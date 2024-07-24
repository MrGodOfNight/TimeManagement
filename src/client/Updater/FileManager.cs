using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace Updater
{
    internal class FileManager
    {
        public int CountOfFiles { get; private set; }
        private List<FileObserver> downloadStatuses;

        public FileManager()
        {
            downloadStatuses = new List<FileObserver>();
        }

        public async Task CompareAndDownloadFiles(string localFolderPath, string serverFolderPath, IProgress<int> progress, Action<List<FileObserver>> updateStatuses)
        {
            string[] localFiles = Directory.GetFiles(localFolderPath);
            string[] serverFiles = Directory.GetFiles(serverFolderPath);

            // Общее количество файлов
            CountOfFiles = serverFiles.Length;

            // Инициализируем список статусов загрузки файлов
            downloadStatuses = new List<FileObserver>();

            foreach (var serverFile in serverFiles)
            {
                string fileName = Path.GetFileName(serverFile);
                string localFilePath = Path.Combine(localFolderPath, fileName);

                bool fileExistsLocally = Array.Exists(localFiles, f => Path.GetFileName(f) == fileName);

                if (!fileExistsLocally || GetFileSize(serverFile) != GetFileSize(localFilePath))
                {
                    // Создаем статус загрузки для текущего файла
                    var downloadStatus = new FileObserver
                    {
                        FileName = fileName,
                        IsDownloading = true
                    };

                    // Добавляем статус загрузки в список
                    downloadStatuses.Add(downloadStatus);

                    // Увеличиваем общий прогресс и выполняем обновление статусов загрузки
                    progress.Report((int)((float)downloadStatuses.Count / CountOfFiles * 100));
                    updateStatuses.Invoke(downloadStatuses);

                    // Скачиваем файл с сервера
                    await DownloadFileAsync(serverFile, localFilePath);

                    // Помечаем загрузку файла как завершенную
                    downloadStatus.IsDownloading = false;

                    // Увеличиваем общий прогресс и выполняем обновление статусов загрузки
                    progress.Report((int)((float)downloadStatuses.Count / CountOfFiles * 100));
                    updateStatuses.Invoke(downloadStatuses);
                }
            }

            // Рекурсивно вызываем метод для всех подпапок
            string[] localDirectories = Directory.GetDirectories(localFolderPath);
            string[] serverDirectories = Directory.GetDirectories(serverFolderPath);

            foreach (var serverDirectory in serverDirectories)
            {
                string directoryName = Path.GetFileName(serverDirectory);
                string localDirectoryPath = Path.Combine(localFolderPath, directoryName);

                bool directoryExistsLocally = Array.Exists(localDirectories, d => Path.GetFileName(d) == directoryName);

                if (!directoryExistsLocally)
                {
                    // Создаем локальную папку, если она не существует
                    Directory.CreateDirectory(localDirectoryPath);
                }

                await CompareAndDownloadFiles(localDirectoryPath, serverDirectory, progress, updateStatuses);
            }
        }

        private long GetFileSize(string filePath)
        {
            return new FileInfo(filePath).Length;
        }

        private async Task DownloadFileAsync(string sourceUrl, string destinationPath)
        {
            using (var client = new WebClient())
            {
                await client.DownloadFileTaskAsync(sourceUrl, destinationPath);
            }
        }
    }
}
