package cmd

import (
	"healer/storage"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "project-up",
	Short: "up project",
	Long:  "",
	Run: runUpCmd,
}

func init() {
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().StringP("project-name", "p", "", "project name (required)")
	err := upCmd.MarkFlagRequired("project-name")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func runUpCmd(cmd *cobra.Command, args []string) {
	projectName, err := cmd.Flags().GetString("project-name")
	if err != nil {
		log.Fatal(err.Error())
	}
	project, err := storage.ReadProject(projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, command := range project.Up.Commands {
		cmd.Printf("execute %s command... \n", command)
		output, err := exec.Command("/bin/bash", "-c", command).CombinedOutput()
		if err != nil {
			log.Fatal(err.Error())
		}
		cmd.Println(string(output))
	}
}
