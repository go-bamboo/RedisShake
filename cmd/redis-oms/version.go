package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "version相关辅助工具",
	Long:  `打印版本信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("version: %v-%v[%v]\n", Version, Revision, BuildDate)
		return nil
	},
}
