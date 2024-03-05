package compose

import (
	"online-learning-platform/pkg/logger"
	"os"
	"os/exec"
)

func StartDockerComposeService() error {
	logger := logger.GetLogger()

	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	logger.Info("Trying to launch the Docker Compose file...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func StopDockerComposeService() error {
	logger := logger.GetLogger()

	cmd := exec.Command("docker-compose", "stop")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	logger.Info("Trying to stop the Docker Compose file...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
