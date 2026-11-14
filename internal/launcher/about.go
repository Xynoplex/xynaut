package launcher

import (
	"fmt"
	"strings"
)

const aboutUsage = `Usage: xynaut about

Show Xynaut's version, source links, image registry, and the container
runtimes detected on this machine.
`

// About implements `xynaut about` — the identity card behind the header.
func About(args []string) int {
	u := NewUI(false)
	for _, a := range args {
		switch a {
		case "--help", "-h":
			fmt.Print(aboutUsage)
			return 0
		default:
			u.Fail("Unknown argument: %s", a)
			return 1
		}
	}

	fmt.Println(u.Bold("Xynoplex · xynaut"))
	fmt.Println()
	fmt.Println("Xynoplex's command-line tool for running and managing local Semiont stacks.")
	fmt.Println()

	runtimes := "none found — install Apple container, Docker, or Podman"
	if found := installedRuntimes(); len(found) > 0 {
		runtimes = strings.Join(found, ", ") + " " + u.Dim("(detected on PATH)")
	}
	row := func(label, value string) {
		fmt.Printf("  %s %s\n", u.Bold(fmt.Sprintf("%-9s", label)), value)
	}
	row("Version", fmt.Sprintf("%s %s", BuildVersion, u.Dim(fmt.Sprintf("(commit %s, built %s)", BuildCommit, BuildDate))))
	row("Source", "https://github.com/Xynoplex/cli")
	row("Issues", "https://github.com/Xynoplex/cli/issues")
	row("Images", imageRegistry+" "+u.Dim("(each knowledge base's documents name the tags it runs)"))
	row("License", "Apache-2.0")
	row("Runtimes", runtimes)

	fmt.Println()
	fmt.Println(u.Dim("This launcher runs the local Semiont stack — graph (Neo4j), vectors"))
	fmt.Println(u.Dim("(Qdrant), inference (Ollama), identity-store (PostgreSQL), traces"))
	fmt.Println(u.Dim("(Jaeger), and the Semiont gateway, worker, smelter, weaver, and Browser —"))
	fmt.Println(u.Dim("by driving your container runtime directly. Try: xynaut start --help"))
	return 0
}
