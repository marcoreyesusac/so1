#include <sys/sysinfo.h>

int main() {
  struct sysinfo info;
  int total_ram, used_ram, free_ram, percentage, carnet;

  // Obtenemos la información de la memoria RAM
  sysinfo(&info);

  // Obtenemos la memoria RAM total
  total_ram = info.totalram;

  // Obtenemos la memoria RAM en uso
  used_ram = info.totalram - info.freeram;

  // Obtenemos la memoria RAM libre
  free_ram = info.freeram;

  // Calculamos el porcentaje de uso
  percentage = (used_ram * 100) / total_ram;

  carnet = 201708850;

  // Escribimos la información en el archivo /proc/ram
  FILE *fp = fopen("/proc/ram", "w");
  if (fp != NULL) {
    fprintf(fp, "%d %d %d %d %d\n", total_ram, used_ram, free_ram, percentage, carnet);
    fclose(fp);
  }

  return 0;
}
