package freebsd

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/otterlabs/jaildeck/internal/domain"
)

const (
	defaultJailConf    = "/etc/jail.conf"
	defaultJailConfDir = defaultJailConf + ".d"
)

var assignmentRE = regexp.MustCompile(`^\s*([a-zA-Z0-9_.-]+)\s*=\s*"?([^";]+)"?\s*;`)

func listJailsFromConfDir(dir string) ([]domain.Jail, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.conf"))
	if err != nil {
		return nil, err
	}

	jails := make([]domain.Jail, 0, len(files))
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		jail, ok := parseSingleJailConf(string(content))
		if !ok {
			continue
		}
		jails = append(jails, jail)
	}

	return jails, nil
}

func parseSingleJailConf(content string) (domain.Jail, bool) {
	lines := strings.Split(content, "\n")

	var jail domain.Jail
	var name string
	values := map[string]string{}

	for _, line := range lines {
		line = stripComment(line)
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasSuffix(line, "{") {
			name = strings.TrimSpace(strings.TrimSuffix(line, "{"))
			continue
		}

		matches := assignmentRE.FindStringSubmatch(line)
		if len(matches) == 3 {
			key := matches[1]
			value := resolveVars(matches[2], name)
			values[key] = value
		}
	}

	if name == "" {
		return domain.Jail{}, false
	}

	jail = domain.Jail{
		Name:     name,
		Status:   domain.JailStatusStopped,
		Hostname: values["host.hostname"],
		IP:       values["ip4"],
		Path:     values["path"],
	}

	if jail.Hostname == "" {
		jail.Hostname = name
	}

	return jail, true
}

func stripComment(line string) string {
	if idx := strings.Index(line, "#"); idx >= 0 {
		return line[:idx]
	}
	return line
}

func resolveVars(value, name string) string {
	value = strings.TrimSpace(value)
	value = strings.Trim(value, `"`)
	value = strings.ReplaceAll(value, "${name}", name)
	return value
}
