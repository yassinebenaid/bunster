package ir

import (
	"fmt"
)

type Compare struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c Compare) togo() string {
	return fmt.Sprintf(
		`if %s %s %s {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.Left.togo(), c.Operator, c.Right.togo())
}

type CompareArithmetics struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c CompareArithmetics) togo() string {
	return fmt.Sprintf(
		`if runtime.NumberCompare(%s, %q, %s) {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.Left.togo(), c.Operator, c.Right.togo())
}

type TestFilesHaveSameDevAndInoNumbers struct {
	File1 Instruction
	File2 Instruction
}

func (c TestFilesHaveSameDevAndInoNumbers) togo() string {
	return fmt.Sprintf(
		`if runtime.FilesHaveSameDevAndIno(%s, %s) {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.File1.togo(), c.File2.togo())
}

type FileIsOlderThan struct {
	File1 Instruction
	File2 Instruction
}

func (c FileIsOlderThan) togo() string {
	return fmt.Sprintf(
		`if runtime.FileIsOlderThan(%s, %s) {
			shell.ExitCode = 0 
		} else {
			shell.ExitCode = 1
		}
		`, c.File1.togo(), c.File2.togo())
}

type TestAgainsStringLength struct {
	String Instruction
	Zero   bool
}

func (c TestAgainsStringLength) togo() string {
	operator := "=="
	if c.Zero {
		operator = "!="
	}

	return fmt.Sprintf(
		`if len(%s) %s 0 {
			shell.ExitCode = 1 
		} else {
			shell.ExitCode = 0
		}
		`, c.String.togo(), operator)
}

type TestFileExists struct {
	File Instruction
}

func (c TestFileExists) togo() string {
	return fmt.Sprintf(
		`if runtime.FileExists(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestDirectoryExists struct {
	File Instruction
}

func (c TestDirectoryExists) togo() string {
	return fmt.Sprintf(
		`if runtime.DirectoryExists(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestBlockSpecialFileExists struct {
	File Instruction
}

func (c TestBlockSpecialFileExists) togo() string {
	return fmt.Sprintf(
		`if runtime.BlockSpecialFileExists(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestCharacterSpecialFileExists struct {
	File Instruction
}

func (c TestCharacterSpecialFileExists) togo() string {
	return fmt.Sprintf(
		`if runtime.CharacterSpecialFileExists(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestRegularFileExists struct {
	File Instruction
}

func (c TestRegularFileExists) togo() string {
	return fmt.Sprintf(
		`if runtime.RegularFileExists(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestFileSGIDIsSet struct {
	File Instruction
}

func (c TestFileSGIDIsSet) togo() string {
	return fmt.Sprintf(
		`if runtime.FileSGIDIsSet(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestFileIsSymbolic struct {
	File Instruction
}

func (c TestFileIsSymbolic) togo() string {
	return fmt.Sprintf(
		`if runtime.FileIsSymbolic(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestFileIsSticky struct {
	File Instruction
}

func (c TestFileIsSticky) togo() string {
	return fmt.Sprintf(
		`if runtime.FileIsSticky(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestFileIsFIFO struct {
	File Instruction
}

func (c TestFileIsFIFO) togo() string {
	return fmt.Sprintf(
		`if runtime.FileIsFIFO(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}

type TestFileIsReadable struct {
	File Instruction
}

func (c TestFileIsReadable) togo() string {
	return fmt.Sprintf(
		`if runtime.FileIsReadable(%s) {
			shell.ExitCode = 0
		} else {
			shell.ExitCode = 1
		}
		`, c.File.togo())
}
