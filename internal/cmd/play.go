package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tken2039/montyhall/internal/montyhall"
	"github.com/tken2039/montyhall/internal/util"
)

type MontyHallOption struct {
	DoorCount  int `validate:"min=3,max=100000"`
	TryCount   int `validate:"min=1,max=1000000"`
	WillChange bool

	DetailMode bool
}

const (
	defaultDoorCount  int  = 3
	defaultTryCount   int  = 100000
	defaultWillChange bool = true
	defaultDetailMode bool = false
)

func NewPlayCmd() *cobra.Command {
	opt := &MontyHallOption{}

	playCmd := &cobra.Command{
		Use:   "play",
		Short: "play initiates the Monty Hall problem",
		Run: func(cmd *cobra.Command, args []string) {
			runPlay(opt)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return util.ValidateParams(*opt)
		},
	}

	playCmd.Flags().IntVar(&opt.DoorCount, "doors", defaultDoorCount, "Set doors to the number of doors.")
	playCmd.Flags().IntVar(&opt.TryCount, "try-count", defaultTryCount, "Set n to the number of games you want to verify.")
	playCmd.Flags().BoolVar(&opt.WillChange, "change", defaultWillChange, "Change determines whether the challenger changes doors midway.")

	playCmd.Flags().BoolVar(&opt.DetailMode, "detail", defaultDetailMode, "Output work details")

	return playCmd
}

func runPlay(opts *MontyHallOption) error {
	v := montyhall.NewValification(opts.DoorCount, opts.TryCount, opts.WillChange, opts.DetailMode)

	return v.Start()
}
