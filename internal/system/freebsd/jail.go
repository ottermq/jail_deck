package freebsd

import (
	"context"
	"fmt"
	"strings"

	"github.com/otterlabs/jaildeck/internal/domain"
	"github.com/otterlabs/jaildeck/internal/system"
)

type jlsOutput struct {
	JailInformation struct {
		Jails []jlsJail `json:"jail"`
	} `json:"jail-information"`
}

type jlsJail struct {
	JID      string   `json:"jid"`
	Name     string   `json:"name"`
	Hostname string   `json:"host.hostname"`
	IP4      []string `json:"ip4.addr"`
	Path     string   `json:"path"`
}

var jlsListCommand = system.Command{
	Name: "jls",
	Args: []string{"--libxo=json", "jid", "name", "host.hostname", "ip4.addr", "path"},
}

func (a *Adapter) List(ctx context.Context) ([]domain.Jail, error) {
	configured, err := a.listConfiguredJails()
	if err != nil {
		return nil, err
	}

	running, err := a.runningJails(ctx)
	if err != nil {
		return nil, err
	}

	return mergeJails(configured, running), nil
}

func (a *Adapter) runService(ctx context.Context, name, action string) (domain.Jail, error) {
	cmd := system.Command{
		Name: "service",
		Args: []string{"jail", action, name},
	}
	fmt.Printf("\ncmd: %v\n", cmd)
	result, err := a.runner.Run(ctx, cmd)

	if err != nil {
		return domain.Jail{}, fmt.Errorf("%q jail %q: %w", action, name, err)
	}

	if found := strings.Contains(result.Stdout, "cannot"); found {
		return domain.Jail{}, fmt.Errorf(
			"%s jail %q failed: %w; exit=%d stdout=%q stderr=%q",
			action,
			name,
			err,
			result.ExitCode,
			result.Stdout,
			result.Stderr,
		)
	}
	err = parseError(name, result.Stdout)
	if err != nil {
		return domain.Jail{}, err
	}

	return a.getJailByName(ctx, name)
}

func parseError(jail, output string) error {
	fmt.Println(output)
	cutted, _ := strings.CutPrefix(output, "Starting jails:")
	outputs := strings.Split(cutted, "\n")
	fmt.Printf("%v", outputs)
	var errStr string
	if len(outputs) > 0 && strings.Contains(outputs[0], "cannot") {
		fmt.Println("found --cannot--")
		var errs []string
		for _, o := range outputs {
			o = strings.TrimSpace(o)
			o, _ = strings.CutPrefix(o, fmt.Sprintf("jail: %s:", jail))
			if o != "" {
				errs = append(errs, o)
			}
		}
		errStr = strings.Join(errs, "/n ")
		return fmt.Errorf("%s %s", outputs[0], errStr)
	}
	return nil
}

func (a *Adapter) Start(ctx context.Context, name string) (domain.Jail, error) {
	return a.runService(ctx, name, "start")
}

func (a *Adapter) Stop(ctx context.Context, name string) (domain.Jail, error) {
	return a.runService(ctx, name, "stop")
}

func (a *Adapter) Restart(ctx context.Context, name string) (domain.Jail, error) {
	return a.runService(ctx, name, "restart")
}
