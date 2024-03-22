package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// var files []string
// var maxDepth int = 0
// var comand string

// // var rootDir string = "/home/airp0wer/go-renamer/"

// func sourceDir() string {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	exPath := filepath.Dir(ex)
// 	fmt.Println(exPath)
// 	return exPath
// }

// type stringFlag struct {
// 	set   bool
// 	value string
// }

// func (sf *stringFlag) Set(x string) error {
// 	sf.value = x
// 	sf.set = true
// 	return nil
// }

// func (sf *stringFlag) String() string {
// 	return sf.value
// }

// var filename stringFlag
// var path stringFlag

// func init() {
// 	flag.Var(&filename, "filename", "the filename")
// 	flag.Var(&path, "path", "the path")
// }

func main() {
	app := &cli.App{
		Name:  "Renamer",
		Usage: "Rename folders and files from cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
			// fmt.Println("Welcome to Renamer!")
			// fmt.Println("use -h for more information")

			// fmt.Printf("Hello %q\n", cCtx.Args().Get(0))

			// return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// rootDir := sourceDir()
	// fmt.Printf("Welcome to Renamer\n=========================\n")
	// flag.Parse()

	// if flag.NFlag() == 0 {
	// 	fmt.Println("use -h or --help")
	// 	os.Exit(0)
	// }

	// if !filename.set {
	// 	fmt.Println("--filename not set")
	// } else {
	// 	fmt.Printf("--filename set to %q\n", filename.value)
	// }

	// fmt.Printf("Chose renamer root folder\n")
	// fmt.Printf("e - %v\n", rootDir)
	// fmt.Printf("o - Type own path\n")
	// fmt.Printf("q - quit\n")
	// var count int
	// var num int
	// flag.IntVar(&count, "x", 1234, "help message for flag x")
	// flag.IntVar(&num, "n", 444, "help message for flag n")

	// flag.Parse()
	// log.Println(flag.)
	// fmt.Println(count)

	// fmt.Scan(&comand)

	// // err := os.File(*os.NewFile(600, "sss.txt"))

	// err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {

	// 	if err != nil {
	// 		log.Println(err)
	// 		// handle possible path err, just in case...
	// 		return err
	// 	}
	// 	if d.IsDir() && strings.Count(path, string(os.PathSeparator))-strings.Count(rootDir, string(os.PathSeparator)) > maxDepth {
	// 		fmt.Println("skip", path)
	// 		return fs.SkipDir
	// 	}

	// 	if d.IsDir() && path != rootDir {
	// 		files = append(files, path)

	// 	}
	// 	// os.Rename()
	// 	return err
	// })

	// if err != nil {
	// 	log.Fatal("error")
	// }

	// fmt.Printf("Files: %v ", len(files))

	// for i, v := range files {
	// 	fmt.Println(v)
	// 	err := os.Rename(v, v+strconv.Itoa(i))
	// 	if err != nil {
	// 		log.Fatal("Path error")
	// 	}
	// }
	// Choose rename
	// - Folders
	// - Files
	// - recursive

	// Chose source

	// Whant to rename another source
	// fmt.Println("Good buy!")
}
