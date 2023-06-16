/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"
	"wwchacalww/go-cem304/adapters/db"
	"wwchacalww/go-cem304/usecase/repository"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var connStr = "postgres://postgres:postgres@db:5432/cem304?sslmode=disable"
var drv, _ = sql.Open("postgres", connStr)
var authDb = db.NewAuthDB(drv)
var userDb = db.NewUserDB(drv)
var classDb = db.NewClassroomDB(drv)
var studentDb = db.NewStudentDB(drv)
var parentDb = db.NewParentDB(drv)
var teacherDb = db.NewTeacherDB(drv)

var authRepository = repository.AuthRepository{Persistence: authDb}
var userRepository = repository.UserRepository{Persistence: userDb}
var classroomRepository = repository.ClassroomRepository{Persistence: classDb}
var studentRepository = repository.StudentRepository{Persistence: studentDb}
var parentRepository = repository.ParentRepository{Persistence: parentDb}
var teacherRepository = repository.TeacherRepository{Persistence: teacherDb}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cem304",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cem304.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
