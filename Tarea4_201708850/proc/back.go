func ReadRamInfo() (int, int, int, int) {
	// Abrimos el archivo /proc/ram
	fp, err := os.Open("/proc/ram")
	if err != nil {
	  return -1, -1, -1, -1
	}
  
	// Leemos la información del archivo
	buf := make([]byte, 128)
	n, err := fp.Read(buf)
	if err != nil {
	  fp.Close()
	  return -1, -1, -1, -1
	}
	fp.Close()
  
	// Separamos la información
	var total_ram, used_ram, free_ram, percentage int
	fmt.Sscanf(string(buf), "%d %d %d %d", &total_ram, &used_ram, &free_ram, &percentage)
  
	return total_ram, used_ram, free_ram, percentage
  }
  