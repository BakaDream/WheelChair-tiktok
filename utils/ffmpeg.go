package utils

import (
	"os/exec"
)

func GetFirstFrame(videoPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-f", "image2", "-t", "0.001", "-y", outputPath)
	err := cmd.Run()
	return err
	//output, err := cmd.CombinedOutput()
	//fmt.Println(string(output))
	//if err != nil {
	//	return fmt.Errorf("ffmpeg error: %s\nOutput: %s", err, output)
	//}
	//return nil
}
