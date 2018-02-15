package garden

import (
	"errors"
	"io"
)

func GardenContainers(gardenAddr string, gardenNetwork string, raw bool, out io.Writer) error {
	return errors.New("GardenContainers: not implemented")
}

/*
type ContainerInfo struct {
	Handle  string
	Info    garden.ContainerInfo
	Metrics garden.Metrics
}

func GardenContainers(gardenAddr string, gardenNetwork string, raw bool, out io.Writer) error {
	client := client.New(connection.New(gardenNetwork, gardenAddr))
	containers, err := client.Containers(nil)
	if err != nil {
		return err
	}

	workPool, err := workpool.NewWorkPool(32)
	if err != nil {
		return err
	}

	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(containers))

	containerInfos := []ContainerInfo{}
	for _, container := range containers {
		container := container
		workPool.Submit(func() {
			defer wg.Done()
			info, err := container.Info()
			if err != nil {
				say.Println(1, say.Red("Failed to fetch container info: %s\n", container.Handle()))
				return
			}
			metrics, err := container.Metrics()
			if err != nil {
				say.Println(1, say.Red("Failed to fetch container metrics: %s\n", container.Handle()))
				return
			}

			lock.Lock()
			defer lock.Unlock()
			containerInfos = append(containerInfos, ContainerInfo{
				container.Handle(),
				info,
				metrics,
			})
		})
	}
	wg.Wait()

	if raw {
		encoded, err := json.MarshalIndent(containerInfos, "", "  ")

		if err != nil {
			return err
		}

		out.Write(encoded)
		return nil
	}

	if len(containerInfos) == 0 {
		say.Println(0, say.Red("No Containers"))
	}
	for _, containerInfo := range containerInfos {
		printContainer(out, containerInfo)
	}
	return nil
}

func printContainer(out io.Writer, containerInfo ContainerInfo) {
	info := containerInfo.Info
	metrics := containerInfo.Metrics
	say.Fprintln(out, 0,
		"%s - %s @ %s",
		say.Green(containerInfo.Handle),
		info.State,
		info.ContainerPath,
	)

	say.Fprintln(out, 1,
		"Memory: %.3f MB",
		float64(metrics.MemoryStat.TotalRss+metrics.MemoryStat.TotalCache-metrics.MemoryStat.TotalInactiveFile)/1024.0/1024.0,
	)

	say.Fprintln(out, 1,
		"Disk: Total:%.3f MB %d Inodes, Exclusive:%.3f MB %d Inodes",
		float64(metrics.DiskStat.TotalBytesUsed)/1024.0/1024.0,
		metrics.DiskStat.TotalInodesUsed,
		float64(metrics.DiskStat.ExclusiveBytesUsed)/1024.0/1024.0,
		metrics.DiskStat.ExclusiveInodesUsed,
	)

	ports := []string{}
	for _, portMapping := range info.MappedPorts {
		ports = append(ports, fmt.Sprintf("%d:%d", portMapping.HostPort, portMapping.ContainerPort))
	}

	say.Fprintln(out, 1,
		"%s=>%s: %s",
		say.Green(info.HostIP),
		say.Green(containerInfo.Handle),
		strings.Join(ports, ","),
	)

	if len(info.Events) > 0 {
		say.Fprintln(out, 1,
			"Events: %s",
			strings.Join(info.Events, ","),
		)
	}

	if len(info.ProcessIDs) > 0 {
		say.Fprintln(out, 1,
			"Running: %d processes",
			len(info.ProcessIDs),
		)
	}

	if len(info.Properties) > 0 {
		say.Fprintln(out, 1,
			"Properties:",
		)
		for key, value := range info.Properties {
			say.Fprintln(out, 2,
				"%s: %s",
				key, value,
			)
		}
	}
}
*/
