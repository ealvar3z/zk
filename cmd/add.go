package cmd

import (
	"fmt"

	zet "github.com/ealvar3z/zk/internal"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a zet",
	RunE: func(c *cobra.Command, args []string) error {
		content, err := openEditor()
		if err != nil {
			return errors.Wrap(err, "no editor")
		}

		tags, err := parseTags(content)
		if err != nil {
			return errors.Wrap(err, "no tags")
		}

		newZet := zet.New(content, tags)
		if err := newZet.Add(); err != nil {
			return errors.Wrap(err, "could not add note")
		}

		fmt.Printf("Zet add with ID %d\n", newZet.ID)
		return nil
	},
}

func openEditor() (string, error) {
	// TODO: Implement this
	return "", errors.New("DIE: could not open editor")
}

func parseTags(content string) ([]string, error) {
	// TODO: Implement this
	return []string{}, errors.New("DIE: could not open editor")
}
