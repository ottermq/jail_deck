package freebsd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ottermq/jaildeck/internal/domain"
	"github.com/ottermq/jaildeck/internal/system"
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
	result, err := a.runner.Run(ctx, cmd)
	if err == nil {
		err = validateJailActionResult(name, action, result)
	}
	if err != nil {
		return domain.Jail{}, &system.CommandError{
			Command: cmd.Name,
			Args:    cmd.Args,
			Result:  result,
			Err:     err,
		}
	}

	return a.getJailByName(ctx, name)
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

func validateJailActionResult(name, action string, result system.CommandResult) error {
	expectedVerbPrefix := "Starting"
	if action == "stop" || action == "restart" {
		expectedVerbPrefix = "Stopping"
	}
	expectedPrefix := fmt.Sprintf("%s jails: %s", expectedVerbPrefix, name)

	// if does not contains the whole expected means it got errors
	errMsg, ok := strings.CutPrefix(result.Stdout, expectedPrefix)
	if ok && len(errMsg) > 1 {
		parts := strings.Split(errMsg, ":")
		if len(parts) > 1 {
			return errors.New(parts[len(parts)-1])
		}
		return errors.New(errMsg)
	}
	return nil
}
