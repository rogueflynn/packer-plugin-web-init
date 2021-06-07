package web

import (
	"context"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepWebCreate struct {
	DirectoryCreateFailed 	bool
	FileCreateFailed 		bool
}

// Run should execute the purpose of this step
func (s *StepWebCreate) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	
	ui.Say("Creating directory")

    f, err1 := os.Create(config.PathName + "/index.html")

    if err1 != nil {
        state.Put("error", err1)
		s.DirectoryCreateFailed = true
		return multistep.ActionHalt
    }

    defer f.Close()
	ui.Say("Creating index.html")
    _, err2 := f.WriteString(
		"<html>\n" + 
		"<head>\n" +
		"</head>\n" +
		"</head>\n" +
		"<body>\n" +
		"</body>\n" +
		"</html>\n")

    if err2 != nil {
        state.Put("error", err2)
		s.FileCreateFailed = true
		return multistep.ActionHalt
    }

	// Determines that should continue to the next step
	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepWebCreate) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)
	ui := state.Get("ui").(packer.Ui)

	if s.DirectoryCreateFailed {
		ui.Say("Directory Creation failed")
		return
	}

	if s.FileCreateFailed {
		ui.Say("File Creation failed")
		return
	}

	if !cancelled && !halted {
		return
	}

}
