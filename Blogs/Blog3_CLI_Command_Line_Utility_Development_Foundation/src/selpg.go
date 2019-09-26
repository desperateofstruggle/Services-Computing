package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	inFilename string
	printdest  string
	pageLen    int
	pageType   string //  "l" for line-based "f" for "/f"
}

var progname string //prgram name

func usage() {
	fmt.Printf("Usage: selpg [-s start] [-e end] [-l linenums | -f] [-d outdest] [inputfilename]\n\n")
	fmt.Printf("Notices:\n")
	fmt.Printf("  -s Number  : start page <Number>.\n")
	fmt.Printf("  -e Number  : end Page <Number>.\n")
	fmt.Printf("  -l Number  : <Number> lines each page. (Default : 72)\n")
	fmt.Printf("  -f         : char '\\f' as the page cut. (Conflict with -l)\n")
	fmt.Printf("  -d Command : <Command> the pipe destination command.(Default : \"\")\n")
	fmt.Printf("  infilename : input file. （Default : stdin. Ps : press Ctrl-D for end)\n\n")
}

func argsInit(args *selpgArgs) {
	pflag.Usage = usage
	pflag.IntVarP(&args.startPage, "s", "s", -1, "Start Page")
	pflag.IntVarP(&args.endPage, "e", "e", -1, "End Page")
	pflag.IntVarP(&args.pageLen, "lop", "l", -0x7f7f7f7f, "Line number of each page")
	pflag.StringVarP(&args.printdest, "dop", "d", "", "the destination of the printer")
	f := pflag.BoolP("xxf", "f", false, "Line cutoff methods")
	pflag.Parse()

	/* check if the -f and -l appear together and assign */
	if *f && args.pageLen == -0x7f7f7f7f {
		args.pageType = "f"
		args.pageLen = -1
	} else if *f && args.pageLen != -0x7f7f7f7f {
		fmt.Fprintf(os.Stderr, "%s: -f and -l cannnot be the input at the same time\n", progname)
		usage()
		os.Exit(1)
	} else if !*f && args.pageLen == -0x7f7f7f7f {
		args.pageType = "l"
		args.pageLen = 72
	} else {
		args.pageType = "l"
	}

	/* set for the file input */
	args.inFilename = ""
	if pflag.NArg() == 1 {
		args.inFilename = string(pflag.Arg(0))
	} else {
		args.inFilename = ""
	}

	/* check the validation of the start page and the end page */
	if args.startPage > args.endPage || args.endPage < 1 {
		fmt.Fprintf(os.Stderr, "Invalid page length:\tstartpage->%d\tendpage->%d\n", args.startPage, args.endPage)
		usage()
		os.Exit(1)
	}
}

func processInput(args *selpgArgs) {
	fin := os.Stdin
	fout := os.Stdout
	cLine := 0
	cPage := 1
	var pipe io.WriteCloser
	var erro error

	/* judge the inputs types */
	if args.inFilename == "" {

	} else {
		fin, erro = os.Open(args.inFilename)
		if erro != nil {
			fmt.Println(erro)
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, args.inFilename)
			usage()
			os.Exit(1)
		}
		defer fin.Close()
	}

	/** judege th outputs types
		simulate the printer through the pipe--->output to the monitor
	**/
	if args.printdest != "" {
		cmd := exec.Command("grep", "-v")
		pipe, erro = cmd.StdinPipe()
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
		defer pipe.Close()
		cmd.Stdout = fout
		cmd.Start()
	}

	/* page distinguish */
	if args.pageType == "l" {
		str := bufio.NewScanner(fin)
		for str.Scan() {
			if cPage >= args.startPage && cPage <= args.endPage {
				fout.Write([]byte(str.Text() + "\n"))
				if args.printdest != "" {
					pipe.Write([]byte(str.Text() + "\n"))
				}
			}
			cLine++
			if cLine >= args.pageLen {
				if cPage >= args.startPage && cPage <= args.endPage {
					fmt.Println("Page ", cPage, " finished")
				}
				cPage++
				cLine = 0
			}
		}
		if cPage < args.endPage {
			fmt.Fprintf(os.Stderr, "%s: could not attach the end page %d\n", progname, args.endPage)
			os.Exit(1)
		} else {
			fmt.Println("\nPage ", "finished", "-->all finished")
		}

	} else if args.pageType == "f" {
		str := bufio.NewReader(fin)
		for {
			p, e := str.ReadString('\f')
			if e == io.EOF {
				fout.Write([]byte(p))
				if cPage >= args.startPage && cPage < args.endPage {
					fmt.Fprintf(os.Stderr, "%s: could not attach the end page %d\n", progname, args.endPage)
					os.Exit(1)
				}
				if cPage < args.startPage {
					fmt.Fprintf(os.Stderr, "%s: could not attach the start page %d\n", progname, args.startPage)
					os.Exit(1)
				}
				fmt.Println("\nPage finished", "-->all finished")
				break
			}
			p = strings.Replace(p, "\f", "\\f\n", -1)

			if cPage >= args.startPage && cPage <= args.endPage {
				fout.Write([]byte(p))
				fmt.Println("Page ", cPage, " finished")
			}
			cPage++
		}
	}
}

func main() {
	progname = os.Args[0]
	var args selpgArgs
	argsInit(&args)
	processInput(&args)
}
