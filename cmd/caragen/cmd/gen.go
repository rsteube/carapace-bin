package cmd

import ()

//func writeFiles() {
//	macros()
//
//	names, descriptions := readCompleters()
//
//	imports := make([]string, 0, len(names))
//	for _, name := range names {
//		imports = append(imports, fmt.Sprintf(`	%v "github.com/rsteube/carapace-bin/completers/%v_completer/cmd"`, varName(name), name))
//	}
//
//	formattedNames := make([]string, 0)
//	for _, name := range names {
//		formattedNames = append(formattedNames, fmt.Sprintf("\t\"%v\",", name))
//	}
//
//	maxlen := 0
//	for _, name := range names {
//		if l := len(name); l > maxlen {
//			maxlen = l
//		}
//	}
//
//	formattedDescriptions := make([]string, 0)
//	for _, name := range names {
//		formattedDescriptions = append(formattedDescriptions, fmt.Sprintf(`	%--`+strconv.Itoa(maxlen+4)+`v"%v",`, fmt.Sprintf(`"%v": `, name), descriptions[name]))
//	}
//
//	cases := make([]string, 0)
//	for _, name := range names {
//		cases = append(cases,
//			fmt.Sprintf(`	case "%v":`, name),
//			fmt.Sprintf(`		%v.Execute()`, varName(name)),
//		)
//	}
//
//	content := fmt.Sprintf(`package cmd

//import (
//%v)
//
//func executeCompleter(completer string) {
//	switch completer {
//%v	}
//}
//`,
//		fmt.Sprintln(strings.Join(imports, "\n")),
//		fmt.Sprintln(strings.Join(cases, "\n")),
//	)
//
//	root, err := rootDir()
//	if err != nil {
//		fmt.Fprintln(os.Stderr, err.Error())
//		os.Exit(1)
//	}
//
//	os.Mkdir(root+"/cmd/carapace/cmd/completers", 0755)
//	os.WriteFile(root+"/cmd/carapace/cmd/completers.go", []byte("//go:build !release\n\n"+content), 0644)
//	os.WriteFile(root+"/cmd/carapace/cmd/completers/name.go", []byte(fmt.Sprintf("package completers\n\nvar names = []string{\n%v\n}\n", strings.Join(formattedNames, "\n"))), 0644)
//	os.WriteFile(root+"/cmd/carapace/cmd/completers/description.go", []byte(fmt.Sprintf("package completers\n\nvar descriptions = map[string]string{\n%v\n}\n", strings.Join(formattedDescriptions, "\n"))), 0644)
//	os.WriteFile(root+"/cmd/carapace/cmd/completers_release.go", []byte("//go:build release\n\n"+strings.Replace(content, "/completers/", "/completers_release/", -1)), 0644)
//	os.RemoveAll(root + "/completers_release")
//	execabs.Command("cp", "-r", root+"/completers", root+"/completers_release").Run()
//
//	for _, name := range names {
//		files, err := os.ReadDir(fmt.Sprintf("%v/completers_release/%v_completer/cmd/", root, name))
//		if err == nil {
//			initFuncs := make([]string, 0)
//			for _, file := range files {
//				if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
//					path := fmt.Sprintf("%v/completers_release/%v_completer/cmd/%v", root, name, file.Name())
//					content, err := os.ReadFile(path)
//					if err == nil && strings.Contains(string(content), "func init() {") {
//						patched := strings.Replace(string(content), "func init() {", fmt.Sprintf("func init_%v() {", strings.TrimSuffix(file.Name(), ".go")), 1)
//						os.WriteFile(path, []byte(patched), os.ModePerm)
//						initFuncs = append(initFuncs, fmt.Sprintf("	init_%v()", strings.TrimSuffix(file.Name(), ".go")))
//					}
//				}
//			}
//
//			path := fmt.Sprintf("%v/completers_release/%v_completer/cmd/root.go", root, name)
//			content, err := os.ReadFile(path)
//			if err == nil {
//				patched := make([]string, 0)
//				for _, line := range strings.Split(string(content), "\n") {
//					patched = append(patched, line)
//					if line == "func Execute() error {" {
//						patched = append(patched, initFuncs...)
//					}
//				}
//				os.WriteFile(path, []byte(strings.Join(patched, "\n")), os.ModePerm)
//			}
//		}
//	}
//}
//
//func varName(name string) string {
//	if name == "go" {
//		return "_go"
//	}
//	return strings.NewReplacer(
//		"-", "_",
//		".", "_",
//	).Replace(name)
//}
//
//func macros() {
//	root, err := rootDir()
//	if err != nil {
//		panic(err.Error)
//	}
//
//	imports := make(map[string]bool)
//	macros := make([]string, 0)
//	descriptions := make(map[string]string)
//
//	r := regexp.MustCompile(`^func Action(?P<name>[^(]+)\((?P<arg>[^(]*)\) carapace.Action {$`)
//	filepath.WalkDir(root+"/pkg/actions", func(path string, d fs.DirEntry, err error) error {
//		if !d.IsDir() && strings.HasSuffix(path, ".go") {
//			file, err := os.Open(path)
//			if err != nil {
//				log.Fatal(err)
//			}
//			defer file.Close()
//
//			pkg := strings.Replace(filepath.Dir(strings.TrimPrefix(path, root+"/pkg/actions/")), "/", ".", -1)
//			_import := fmt.Sprintf(`	%v "github.com/rsteube/carapace-bin/pkg/actions/%v"`, strings.Replace(pkg, ".", "_", -1), strings.Replace(pkg, ".", "/", -1))
//
//			scanner := bufio.NewScanner(file)
//			for scanner.Scan() {
//				if t := scanner.Text(); strings.HasPrefix(t, "func Action") {
//					if r.MatchString(t) {
//						matches := r.FindStringSubmatch(t)
//
//						_func := fmt.Sprintf("%v.Action%v", strings.Replace(pkg, ".", "_", -1), matches[1])
//
//						disableflagParsing := ""
//						if pkg == "bridge" {
//							disableflagParsing = ".NoFlag()"
//						}
//
//						if arg := matches[2]; strings.Contains(arg, ",") {
//							macros = append(macros, "// TODO unsupported signature: "+t)
//							continue
//						} else if arg == "" {
//							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroN(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
//						} else if strings.Contains(arg, "...") {
//							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroV(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
//						} else {
//							macros = append(macros, fmt.Sprintf(`"%v.%v": spec.MacroI(%v)%v,`, pkg, matches[1], _func, disableflagParsing))
//						}
//						imports[_import] = true
//					}
//				} else if strings.HasPrefix(t, "// Action") {
//					if splitted := strings.SplitN(strings.TrimPrefix(t, "// Action"), " ", 2); len(splitted) > 1 {
//						descriptions[pkg+"."+splitted[0]] = splitted[1]
//					}
//				}
//			}
//
//		}
//		return nil
//	})
//
//	sortedImports := make([]string, 0)
//	for i := range imports {
//		sortedImports = append(sortedImports, i)
//	}
//	sort.Strings(sortedImports)
//
//	sortedDescriptions := make([]string, 0)
//	for key, value := range descriptions {
//		sortedDescriptions = append(sortedDescriptions, fmt.Sprintf(`"%v": "%v",`, key, value))
//	}
//	sort.Strings(sortedDescriptions)
//
//	content := fmt.Sprintf(`package cmd
//
//import (
//%v
//	spec "github.com/rsteube/carapace-spec"
//)
//
//var macros = map[string]spec.Macro{
//%v
//}
//
//var macroDescriptions = map[string]string {
//%v
//}
//`, strings.Join(sortedImports, "\n"), strings.Join(macros, "\n"), strings.Join(sortedDescriptions, "\n"))
//
//	os.WriteFile(root+"/cmd/carapace/cmd/macros.go", []byte(content), 0644)
//	execabs.Command("go", "fmt", root+"/cmd/carapace/cmd/macros.go").Run()
//
//}
