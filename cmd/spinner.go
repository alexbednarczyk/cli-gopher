package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/briandowns/spinner"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var spinnerCmd = &cobra.Command{
	Use:   "spinner",
	Short: "A simple spinner command",
	Long:  `A simple spinner command that displays a spinner based on user configuration`,
	Args:  cobra.NoArgs,
	RunE:  displaySpinner,
}

func displaySpinner(cm *cobra.Command, args []string) error {
	ticker := time.NewTicker(time.Second)
	quit := make(chan bool)
	expression := "spin"

	switch expression {
	case "spin":
		fallthrough
	default:
		go spin(ticker, quit)
	}

	fmt.Fprintln(cm.OutOrStdout(), "Press Enter to quit")
	_, _ = bufio.NewReader(cm.InOrStdin()).ReadBytes('\n')
	quit <- true
	ticker.Stop()
	return nil
}

func spin(ticker *time.Ticker, quit chan bool) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	const defaultCharset = 9
	s := spinner.New(spinner.CharSets[defaultCharset], 200*time.Millisecond) // Build our new default spinner
	s.Start()                                                                // Start the spinner
	s.FinalMSG = "Done!"
	defer s.Stop()

	if viper.ConfigFileUsed() != "" {
		viper.OnConfigChange(func(e fsnotify.Event) {
			charSetIndex := viper.GetInt("charset")
			if charSetIndex < 0 || charSetIndex >= len(spinner.CharSets) {
				charSetIndex = defaultCharset
			}
			s.UpdateCharSet(spinner.CharSets[charSetIndex])
			s.Restart()
			s.FinalMSG = "Modified Done!"
		})
		viper.WatchConfig()
	}

Loop:
	for {
		select {
		case <-ticker.C:
			time.Sleep(100 * time.Millisecond) // Run for some time to simulate work
		case <-quit:
			break Loop
		}
	}
	ticker.Stop()
}

func init() {
	rootCmd.AddCommand(spinnerCmd)
}
