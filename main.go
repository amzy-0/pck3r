package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	mag    = color.New(color.FgMagenta).SprintFunc()
	reset  = color.New(color.Reset).SprintFunc()
)

func sysErr(msg string) string {
	return fmt.Sprintf("\n%s尸⼕长㇌尺 : ERROR !\n%s%s", red(""), red(msg), reset(""))
}

func sysOk(msg string) string {
	return fmt.Sprintf("\n%s尸⼕长㇌尺 :\n %s%s", green(""), green(msg), reset(""))
}

func printHelp() {
	// Try to read /bin/pck3r-help first
	helpPath := "/bin/pck3r-help"
	content, err := ioutil.ReadFile(helpPath)
	if err != nil {
		// fallback to local README.md starting from line 24
		readmePath := "README.md"
		content, err = ioutil.ReadFile(readmePath)
		if err != nil {
			fmt.Println(sysErr("Help file not found"))
			return
		}
		lines := bytes.Split(content, []byte{'\n'})
		if len(lines) > 24 {
			content = bytes.Join(lines[24:], []byte{'\n'})
		}
	}
	fmt.Println(yellow(string(content)))
}

func afterEmpty(command string, helpContents string) {
	if helpContents == "" {
		helpContents = fmt.Sprintf("$ pck3r %s hello", command)
	}
	fmt.Printf("%s%sAfter \"%s\" is empty!\n%s%s\n", sysErr(""), red(""), command, yellow(helpContents), reset(""))
}

func pkgFind(packageName string) {
	if packageName == "" {
		afterEmpty("pkg", "$ pck3r pkg hello")
		return
	}
	cmd := exec.Command("apt", "search", packageName+".+*")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func sysCommand(sysMigration string) {
	if sysMigration == "" {
		afterEmpty("sys", "$ pck3r sys {update/upgrade/updgr}\nupdgr : update and full-upgrade, packages.")
		return
	}
	var cmd *exec.Cmd
	switch sysMigration {
	case "update":
		cmd = exec.Command("sudo", "apt", "update")
	case "upgrade":
		cmd = exec.Command("sudo", "apt", "upgrade")
	case "updgr":
		cmd = exec.Command("bash", "-c", "sudo apt update && sudo apt -y full-upgrade")
	default:
		fmt.Printf("%s%sInvalid sys command: %s%s\n", sysErr(""), red(""), sysMigration, reset(""))
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func clearCommand() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	fmt.Println(sysOk("This is a funny clear command :D"))
}

func updateCommand() {
	err := os.Chdir("/opt/pck3r")
	if err != nil {
		fmt.Println(sysErr("Failed to change directory to /opt/pck3r"))
		return
	}
	cmd := exec.Command("bash", "-c", "sudo git pull && sudo git restore .")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func installCommand(packageName string) {
	if packageName == "" {
		afterEmpty("install", "$ pck3r install {package name}")
		return
	}
	switch packageName {
	case "nodejs":
		installNodejs()
	case "ohmyzsh":
		installOhMyZsh()
	default:
		handleGenericInstall(packageName)
	}
}

func handleGenericInstall(packageName string) {
	fmt.Printf("%s\n[WAIT FOR PROCESSING]\n%s\n", sysOk(""), yellow(""))
	cmd := exec.Command("sudo", "apt", "install", "-y", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s%sPackage(s) or Command(s) not found: %s%s\n", sysErr(""), red(""), packageName, reset(""))
	}
}

func installNodejs() {
	script := `
echo ` + yellow("") + `
curl -fsSL https://deb.nodesource.com/setup_22.x | sudo -E bash - ;
echo ` + cyan("") + ` ; sudo apt install -y nodejs ;
sudo apt-get update && echo ` + mag("") + ` ;
`
	cmd := exec.Command("bash", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		fmt.Println(sysOk("Node.js installed successfully!"))
	} else {
		fmt.Printf("%s%s\nPlease retry...\n$pck3r install nodejs%s\n", sysErr(""), red(""), reset(""))
	}
}

func installOhMyZsh() {
	fmt.Println("Installing Oh My Zsh...")
	cmdGit := exec.Command("sudo", "apt", "install", "-y", "git")
	cmdGit.Stdout = os.Stdout
	cmdGit.Stderr = os.Stderr
	_ = cmdGit.Run()

	cmdCurlCheck := exec.Command("curl", "--version")
	if err := cmdCurlCheck.Run(); err != nil {
		fmt.Printf("%s%s\"curl\" is required for using \"oh-my-zsh\" ; installing curl...%s\n", sysErr(""), red(""), reset(""))
		cmdCurlInstall := exec.Command("sudo", "apt", "install", "-y", "curl")
		if err := cmdCurlInstall.Run(); err != nil {
			fmt.Printf("%s%sFailed to install curl. Please install it manually.%s\n", sysErr(""), red(""), reset(""))
			return
		}
	}

	installCmd := exec.Command("bash", "-c", `sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err == nil {
		fmt.Println(sysOk("Oh My Zsh installed successfully!"))
	} else {
		fmt.Printf("%s%sOh My Zsh installation failed.%s\n", sysErr(""), red(""), reset(""))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s%sNo command provided. Use \"--help\" for a list of available commands.%s\n", sysErr(""), red(""), reset(""))
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "-h" || arg == "--help" {
		printHelp()
		os.Exit(0)
	}

	validCommands := map[string]bool{
		"clear":     true,
		"update":    true,
		"install":   true,
		"uninstall": true,
		"rm":        true,
		"sys":       true,
		"pkg":       true,
		"version":   true,
	}

	if !validCommands[arg] {
		fmt.Printf("%s%sCommand not found: %s%s\n", sysErr(""), red(""), arg, reset(""))
		os.Exit(1)
	}

	switch arg {
	case "clear":
		clearCommand()
	case "update":
		updateCommand()
	case "install":
		var pkg string
		if len(os.Args) > 2 {
			pkg = os.Args[2]
		}
		installCommand(pkg)
	case "sys":
		var sysArg string
		if len(os.Args) > 2 {
			sysArg = os.Args[2]
		}
		sysCommand(sysArg)
	case "pkg":
		var pkg string
		if len(os.Args) > 2 {
			pkg = os.Args[2]
		}
		pkgFind(pkg)
	case "version":
		fmt.Printf("\b%s\bversion : 1.0\n", sysOk(""))
	default:
		fmt.Printf("%s%sCommand not found: %s%s\n", sysErr(""), red(""), arg, reset(""))
		os.Exit(1)
	}
}
