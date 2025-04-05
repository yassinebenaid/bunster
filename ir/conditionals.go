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
	return fmt.Sprintf("testResult = %s %s %s\n", c.Left.togo(), c.Operator, c.Right.togo())
}

type CompareArithmetics struct {
	Left     Instruction
	Operator string
	Right    Instruction
}

func (c CompareArithmetics) togo() string {
	return fmt.Sprintf("testResult = runtime.NumberCompare(%s, %q, %s)\n", c.Left.togo(), c.Operator, c.Right.togo())
}

type TestFilesHaveSameDevAndInoNumbers struct {
	File1 Instruction
	File2 Instruction
}

func (c TestFilesHaveSameDevAndInoNumbers) togo() string {
	return fmt.Sprintf("testResult = runtime.FilesHaveSameDevAndIno(shell.Path(%s), shell.Path(%s))\n", c.File1.togo(), c.File2.togo())
}

type FileIsOlderThan struct {
	File1 Instruction
	File2 Instruction
}

func (c FileIsOlderThan) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsOlderThan(shell.Path(%s), shell.Path(%s))\n", c.File1.togo(), c.File2.togo())
}

type TestAgainsStringLength struct {
	String Instruction
	Zero   bool
}

func (c TestAgainsStringLength) togo() string {
	operator := "=="
	if !c.Zero {
		operator = "!="
	}

	return fmt.Sprintf("testResult = len(%s) %s 0\n", c.String.togo(), operator)
}

type TestFileExists struct {
	File Instruction
}

func (c TestFileExists) togo() string {
	return fmt.Sprintf("testResult = runtime.FileExists(shell.Path(%s))\n", c.File.togo())
}

type TestDirectoryExists struct {
	File Instruction
}

func (c TestDirectoryExists) togo() string {
	return fmt.Sprintf("testResult = runtime.DirectoryExists(shell.Path(%s))\n", c.File.togo())
}

type TestBlockSpecialFileExists struct {
	File Instruction
}

func (c TestBlockSpecialFileExists) togo() string {
	return fmt.Sprintf("testResult = runtime.BlockSpecialFileExists(shell.Path(%s))\n", c.File.togo())
}

type TestCharacterSpecialFileExists struct {
	File Instruction
}

func (c TestCharacterSpecialFileExists) togo() string {
	return fmt.Sprintf("testResult = runtime.CharacterSpecialFileExists(shell.Path(%s))\n", c.File.togo())
}

type TestRegularFileExists struct {
	File Instruction
}

func (c TestRegularFileExists) togo() string {
	return fmt.Sprintf("testResult = runtime.RegularFileExists(shell.Path(%s))\n", c.File.togo())
}

type TestFileSGIDIsSet struct {
	File Instruction
}

func (c TestFileSGIDIsSet) togo() string {
	return fmt.Sprintf("testResult = runtime.FileSGIDIsSet(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsOwnedByEffectiveGroup struct {
	File Instruction
}

func (c TestFileIsOwnedByEffectiveGroup) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsOwnedByEffectiveGroup(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsOwnedByEffectiveUser struct {
	File Instruction
}

func (c TestFileIsOwnedByEffectiveUser) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsOwnedByEffectiveUser(shell.Path(%s))\n", c.File.togo())
}

type TestFileHasBeenModifiedSinceLastRead struct {
	File Instruction
}

func (c TestFileHasBeenModifiedSinceLastRead) togo() string {
	return fmt.Sprintf("testResult = runtime.FileHasBeenModifiedSinceLastRead(shell.Path(%s))\n", c.File.togo())
}

type TestFileSUIDIsSet struct {
	File Instruction
}

func (c TestFileSUIDIsSet) togo() string {
	return fmt.Sprintf("testResult = runtime.FileSUIDIsSet(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsSymbolic struct {
	File Instruction
}

func (c TestFileIsSymbolic) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsSymbolic(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsSticky struct {
	File Instruction
}

func (c TestFileIsSticky) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsSticky(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsFIFO struct {
	File Instruction
}

func (c TestFileIsFIFO) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsFIFO(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsReadable struct {
	File Instruction
}

func (c TestFileIsReadable) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsReadable(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsWritable struct {
	File Instruction
}

func (c TestFileIsWritable) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsWritable(shell.Path(%s))\n", c.File.togo())
}

type TestFileIsExecutable struct {
	File Instruction
}

func (c TestFileIsExecutable) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsExecutable(shell.Path(%s))\n", c.File.togo())
}

type TestFileHasAPositiveSize struct {
	File Instruction
}

func (c TestFileHasAPositiveSize) togo() string {
	return fmt.Sprintf("testResult = runtime.FileHasAPositiveSize(shell.Path(%s))\n", c.File.togo())
}

type TestFileDescriptorIsTerminal struct {
	File Instruction
}

func (c TestFileDescriptorIsTerminal) togo() string {
	return fmt.Sprintf("testResult = runtime.FileDescriptorIsTerminal(streamManager, %s)\n", c.File.togo())
}

type TestFileIsSocket struct {
	File Instruction
}

func (c TestFileIsSocket) togo() string {
	return fmt.Sprintf("testResult = runtime.FileIsSocket(shell.Path(%s))\n", c.File.togo())
}

type TestVarIsSet struct {
	Name Instruction
}

func (c TestVarIsSet) togo() string {
	return fmt.Sprintf("testResult = shell.VarIsSet(%s)\n", c.Name.togo())
}
