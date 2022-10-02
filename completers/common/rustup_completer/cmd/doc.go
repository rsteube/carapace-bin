package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Open the documentation for the current toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().Bool("alloc", false, "The Rust core allocation and collections library")
	docCmd.Flags().Bool("book", false, "The Rust Programming Language book")
	docCmd.Flags().Bool("cargo", false, "The Cargo Book")
	docCmd.Flags().Bool("core", false, "The Rust Core Library")
	docCmd.Flags().Bool("edition-guide", false, "The Rust Edition Guide")
	docCmd.Flags().Bool("embedded-book", false, "The Embedded Rust Book")
	docCmd.Flags().BoolP("help", "h", false, "Prints help information")
	docCmd.Flags().Bool("nomicon", false, "The Dark Arts of Advanced and Unsafe Rust Programming")
	docCmd.Flags().Bool("path", false, "Only print the path to the documentation")
	docCmd.Flags().Bool("proc_macro", false, "A support library for macro authors when defining new macros")
	docCmd.Flags().Bool("reference", false, "The Rust Reference")
	docCmd.Flags().Bool("rust-by-example", false, "A collection of runnable examples that illustrate various Rust concepts and standard")
	docCmd.Flags().Bool("rustc", false, "The compiler for the Rust programming language")
	docCmd.Flags().Bool("rustdoc", false, "Generate documentation for Rust projects")
	docCmd.Flags().Bool("std", false, "Standard library API documentation")
	docCmd.Flags().Bool("test", false, "Support code for rustc's built in unit-test and micro-benchmarking framework")
	docCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	docCmd.Flags().Bool("unstable-book", false, "The Unstable Book")
	rootCmd.AddCommand(docCmd)

	carapace.Gen(docCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
