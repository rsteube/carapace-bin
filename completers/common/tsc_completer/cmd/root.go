package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tsc_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsc",
	Short: "The TypeScript Compiler",
	Long:  "https://www.typescriptlang.org/docs/handbook/compiler-options.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Show all compiler options.")
	rootCmd.Flags().Bool("allowJs", false, "Allow JavaScript files to be a part of your program. Use the `checkJS` option to get errors from these files.")
	rootCmd.Flags().Bool("allowSyntheticDefaultImports", false, "Allow 'import x from y' when a module doesn't have a default export.")
	rootCmd.Flags().Bool("allowUmdGlobalAccess", false, "Allow accessing UMD globals from modules.")
	rootCmd.Flags().Bool("allowUnreachableCode", false, "Disable error reporting for unreachable code.")
	rootCmd.Flags().Bool("allowUnusedLabels", false, "Disable error reporting for unused labels.")
	rootCmd.Flags().Bool("alwaysStrict", false, "Ensure 'use strict' is always emitted.")
	rootCmd.Flags().Bool("assumeChangesOnlyAffectDirectDependencies", false, "Have recompiles in projects that use `incremental` and `watch` mode assume that changes within a file will only affect files directly depending on it.")
	rootCmd.Flags().String("baseUrl", "", "Specify the base directory to resolve non-relative module names.")
	rootCmd.Flags().String("build", "b", "Build one or more projects and their dependencies, if out of date")
	rootCmd.Flags().Bool("checkJs", false, "Enable error reporting in type-checked JavaScript files.")
	rootCmd.Flags().Bool("clean", false, "Delete the outputs of all projects")
	rootCmd.Flags().Bool("composite", false, "Enable constraints that allow a TypeScript project to be used with project references.")
	rootCmd.Flags().String("declaration", "", "Generate .d.ts files from TypeScript and JavaScript files in your project.")
	rootCmd.Flags().String("declarationDir", "", "Specify the output directory for generated declaration files.")
	rootCmd.Flags().Bool("declarationMap", false, "Create sourcemaps for d.ts files.")
	rootCmd.Flags().Bool("diagnostics", false, "Output compiler performance information after building.")
	rootCmd.Flags().Bool("disableReferencedProjectLoad", false, "Reduce the number of projects loaded automatically by TypeScript.")
	rootCmd.Flags().Bool("disableSizeLimit", false, "Remove the 20mb cap on total source code size for JavaScript files in the TypeScript language server.")
	rootCmd.Flags().Bool("disableSolutionSearching", false, "Opt a project out of multi-project reference checking when editing.")
	rootCmd.Flags().Bool("disableSourceOfProjectReferenceRedirect", false, "Disable preferring source files instead of declaration files when referencing composite projects")
	rootCmd.Flags().Bool("downlevelIteration", false, "Emit more compliant, but verbose and less performant JavaScript for iteration.")
	rootCmd.Flags().StringP("dry", "d", "", "Show what would be built (or deleted, if specified with '--clean')")
	rootCmd.Flags().Bool("emitBOM", false, "Emit a UTF-8 Byte Order Mark (BOM) in the beginning of output files.")
	rootCmd.Flags().Bool("emitDeclarationOnly", false, "Only output d.ts files and not JavaScript files.")
	rootCmd.Flags().Bool("emitDecoratorMetadata", false, "Emit design-type metadata for decorated declarations in source files.")
	rootCmd.Flags().Bool("esModuleInterop", false, "Emit additional JavaScript to ease support for importing CommonJS modules. This enables `allowSyntheticDefaultImports` for type compatibility.")
	rootCmd.Flags().Bool("exactOptionalPropertyTypes", false, "Interpret optional property types as written, rather than adding 'undefined'.")
	rootCmd.Flags().String("excludeDirectories", "", "Remove a list of directories from the watch process.")
	rootCmd.Flags().Bool("excludeFiles", false, "Remove a list of files from the watch mode's processing.")
	rootCmd.Flags().Bool("experimentalDecorators", false, "Enable experimental support for TC39 stage 2 draft decorators.")
	rootCmd.Flags().Bool("explainFiles", false, "Print files read during the compilation including why it was included.")
	rootCmd.Flags().Bool("extendedDiagnostics", false, "Output more detailed compiler performance information after building.")
	rootCmd.Flags().String("fallbackPolling", "", "Specify what approach the watcher should use if the system runs out of native file watchers.")
	rootCmd.Flags().StringP("force", "f", "", "Build all projects, including those that appear to be up to date")
	rootCmd.Flags().Bool("forceConsistentCasingInFileNames", false, "Ensure that casing is correct in imports.")
	rootCmd.Flags().String("generateCpuProfile", "", "Emit a v8 CPU profile of the compiler run for debugging.")
	rootCmd.Flags().Bool("generateTrace", false, "Generates an event trace and a list of types.")
	rootCmd.Flags().StringP("help", "h", "", "Print this message.")
	rootCmd.Flags().Bool("importHelpers", false, "Allow importing helper functions from tslib once per project, instead of including them per-file.")
	rootCmd.Flags().Bool("importsNotUsedAsValues", false, "Specify emit/checking behavior for imports that are only used for types")
	rootCmd.Flags().StringP("incremental", "i", "", "Enable incremental compilation")
	rootCmd.Flags().Bool("init", false, "Initializes a TypeScript project and creates a tsconfig.json file.")
	rootCmd.Flags().Bool("inlineSourceMap", false, "Include sourcemap files inside the emitted JavaScript.")
	rootCmd.Flags().Bool("inlineSources", false, "Include source code in the sourcemaps inside the emitted JavaScript.")
	rootCmd.Flags().Bool("isolatedModules", false, "Ensure that each file can be safely transpiled without relying on other imports.")
	rootCmd.Flags().String("jsx", "", "Specify what JSX code is generated.")
	rootCmd.Flags().String("jsxFactory", "", "Specify the JSX factory function used when targeting React JSX emit, e.g. 'React.createElement' or 'h'")
	rootCmd.Flags().String("jsxFragmentFactory", "", "Specify the JSX Fragment reference used for fragments when targeting React JSX emit e.g. 'React.Fragment' or 'Fragment'.")
	rootCmd.Flags().String("jsxImportSource", "", "Specify module specifier used to import the JSX factory functions when using `jsx: react-jsx*`.`")
	rootCmd.Flags().String("lib", "", "Specify a set of bundled library declaration files that describe the target runtime environment.")
	rootCmd.Flags().Bool("listEmittedFiles", false, "Print the names of emitted files after a compilation.")
	rootCmd.Flags().Bool("listFiles", false, "Print all of the files read during the compilation.")
	rootCmd.Flags().Bool("listFilesOnly", false, "Print names of files that are part of the compilation and then stop processing.")
	rootCmd.Flags().String("locale", "", "Set the language of the messaging from TypeScript. This does not affect emit.")
	rootCmd.Flags().String("mapRoot", "", "Specify the location where debugger should locate map files instead of generated locations.")
	rootCmd.Flags().String("maxNodeModuleJsDepth", "", "Specify the maximum folder depth used for checking JavaScript files from `node_modules`. Only applicable with `allowJs`.")
	rootCmd.Flags().StringP("module", "m", "", "VAL  Specify what module code is generated.")
	rootCmd.Flags().String("moduleResolution", "", "Specify how TypeScript looks up a file from a given module specifier.")
	rootCmd.Flags().String("newLine", "", "Set the newline character for emitting files.")
	rootCmd.Flags().Bool("noEmit", false, "Disable emitting files from a compilation.")
	rootCmd.Flags().Bool("noEmitHelpers", false, "Disable generating custom helper functions like `__extends` in compiled output.")
	rootCmd.Flags().Bool("noEmitOnError", false, "Disable emitting files if any type checking errors are reported.")
	rootCmd.Flags().Bool("noErrorTruncation", false, "Disable truncating types in error messages.")
	rootCmd.Flags().Bool("noFallthroughCasesInSwitch", false, "Enable error reporting for fallthrough cases in switch statements.")
	rootCmd.Flags().Bool("noImplicitAny", false, "Enable error reporting for expressions and declarations with an implied `any` type..")
	rootCmd.Flags().Bool("noImplicitOverride", false, "Ensure overriding members in derived classes are marked with an override modifier.")
	rootCmd.Flags().Bool("noImplicitReturns", false, "Enable error reporting for codepaths that do not explicitly return in a function.")
	rootCmd.Flags().Bool("noImplicitThis", false, "Enable error reporting when `this` is given the type `any`.")
	rootCmd.Flags().Bool("noLib", false, "Disable including any library files, including the default lib.d.ts.")
	rootCmd.Flags().Bool("noPropertyAccessFromIndexSignature", false, "Enforces using indexed accessors for keys declared using an indexed type")
	rootCmd.Flags().Bool("noResolve", false, "Disallow `import`s, `require`s or `<reference>`s from expanding the number of files TypeScript should add to a project.")
	rootCmd.Flags().Bool("noUncheckedIndexedAccess", false, "Include 'undefined' in index signature results")
	rootCmd.Flags().Bool("noUnusedLocals", false, "Enable error reporting when a local variables aren't read.")
	rootCmd.Flags().Bool("noUnusedParameters", false, "Raise an error when a function parameter isn't read")
	rootCmd.Flags().String("outDir", "", "Specify an output folder for all emitted files.")
	rootCmd.Flags().String("outFile", "", "Specify a file that bundles all outputs into one JavaScript file. If `declaration` is true, also designates a file that bundles all .d.ts output.")
	rootCmd.Flags().Bool("paths", false, "Specify a set of entries that re-map imports to additional lookup locations.")
	rootCmd.Flags().String("plugins", "", "List of language service plugins.")
	rootCmd.Flags().Bool("preserveConstEnums", false, "Disable erasing `const enum` declarations in generated code.")
	rootCmd.Flags().Bool("preserveSymlinks", false, "Disable resolving symlinks to their realpath. This correlates to the same flag in node.")
	rootCmd.Flags().Bool("preserveWatchOutput", false, "Disable wiping the console in watch mode")
	rootCmd.Flags().Bool("pretty", false, "Enable color and formatting in TypeScript's output to make compiler errors easier to read")
	rootCmd.Flags().String("reactNamespace", "", "Specify the object invoked for `createElement`. This only applies when targeting `react` JSX emit.")
	rootCmd.Flags().Bool("removeComments", false, "Disable emitting comments.")
	rootCmd.Flags().Bool("resolveJsonModule", false, "Enable importing .json files")
	rootCmd.Flags().String("rootDir", "", "Specify the root folder within your source files.")
	rootCmd.Flags().String("rootDirs", "", "Allow multiple folders to be treated as one when resolving modules.")
	rootCmd.Flags().Bool("showConfig", false, "Print the final configuration instead of building.")
	rootCmd.Flags().Bool("skipDefaultLibCheck", false, "Skip type checking .d.ts files that are included with TypeScript.")
	rootCmd.Flags().Bool("skipLibCheck", false, "Skip type checking all .d.ts files.")
	rootCmd.Flags().Bool("sourceMap", false, "Create source map files for emitted JavaScript files.")
	rootCmd.Flags().String("sourceRoot", "", "Specify the root path for debuggers to find the reference source code.")
	rootCmd.Flags().Bool("strict", false, "Enable all strict type-checking options.")
	rootCmd.Flags().Bool("strictBindCallApply", false, "Check that the arguments for `bind`, `call`, and `apply` methods match the original function.")
	rootCmd.Flags().Bool("strictFunctionTypes", false, "When assigning functions, check to ensure parameters and the return values are subtype-compatible.")
	rootCmd.Flags().Bool("strictNullChecks", false, "When type checking, take into account `null` and `undefined`.")
	rootCmd.Flags().Bool("strictPropertyInitialization", false, "Check for class properties that are declared but not set in the constructor.")
	rootCmd.Flags().Bool("stripInternal", false, "Disable emitting declarations that have `@internal` in their JSDoc comments.")
	rootCmd.Flags().Bool("synchronousWatchDirectory", false, "Synchronously call callbacks and update the state of directory watchers on platforms that don`t support recursive watching natively.")
	rootCmd.Flags().StringP("target", "t", "", "Set the JavaScript language version for emitted JavaScript and include compatible library declarations")
	rootCmd.Flags().Bool("traceResolution", false, "Log paths used during the `moduleResolution` process.")
	rootCmd.Flags().String("tsBuildInfoFile", "", "Specify the folder for .tsbuildinfo incremental compilation files.")
	rootCmd.Flags().String("typeRoots", "", "Specify multiple folders that act like `./node_modules/@types`.")
	rootCmd.Flags().String("types", "", "Specify type package names to be included without being referenced in a source file.")
	rootCmd.Flags().Bool("useDefineForClassFields", false, "Emit ECMAScript-standard-compliant class fields.")
	rootCmd.Flags().Bool("useUnknownInCatchVariables", false, "Type catch clause variables as 'unknown' instead of 'any'.")
	rootCmd.Flags().String("verbose", "", "Enable verbose logging")
	rootCmd.Flags().StringP("version", "v", "", "Print the compiler's version.")
	rootCmd.Flags().StringP("watch", "w", "", "Watch input files.")
	rootCmd.Flags().String("watchDirectory", "", "Specify how directories are watched on systems that lack recursive file-watching functionality.")
	rootCmd.Flags().String("watchFile", "", "Specify how the TypeScript watch mode works.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"baseUrl":         carapace.ActionDirectories(),
		"declarationDir":  carapace.ActionDirectories(),
		"fallbackPolling": carapace.ActionValues("fixedinterval", "priorityinterval", "dynamicpriority", "fixedchunksize"),
		"jsx":             carapace.ActionValues("preserve", "react-native", "react", "react-jsx", "react-jsxdev"),
		"lib": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLibs().Invoke(c).Filter(c.Parts).ToA()
		}),
		"mapRoot":          carapace.ActionDirectories(),
		"module":           carapace.ActionValues("none", "commonjs", "amd", "system", "umd", "es6", "es2015", "es2020", "esnext"),
		"moduleResolution": carapace.ActionValues("node", "classic"),
		"newLine":          carapace.ActionValues("crlf", "lf"),
		"outDir":           carapace.ActionDirectories(),
		"outFile":          carapace.ActionFiles(),
		"rootDir":          carapace.ActionDirectories(),
		"rootDirs": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionDirectories()
		}),
		"sourceRoot":      carapace.ActionDirectories(),
		"target":          carapace.ActionValues("es3", "es5", "es6", "es2015", "es2016", "es2017", "es2018", "es2019", "es2020", "es2021", "esnext"),
		"tsBuildInfoFile": carapace.ActionFiles(),
		"typeRoots": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionDirectories()
		}),
		"watchDirectory": carapace.ActionValues("usefsevents", "fixedpollinginterval", "dynamicprioritypolling", "fixedchunksizepolling"),
		"watchFile":      carapace.ActionValues("fixedpollinginterval", "prioritypollinginterval", "dynamicprioritypolling", "fixedchunksizepolling", "usefsevents", "usefseventsonparentdirectory"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
