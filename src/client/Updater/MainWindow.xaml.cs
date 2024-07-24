using System.Collections.ObjectModel;
using System.Diagnostics;
using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace Updater
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public ObservableCollection<FileObserver> DownloadStatuses { get; set; }
        public MainWindow()
        {
            InitializeComponent();

            // Инициализация коллекции статусов загрузки
            DownloadStatuses = new ObservableCollection<FileObserver>();

            // Привязка списка файлов к ListBox
            fileListBox.ItemsSource = DownloadStatuses;
        }

        private void cancel_Click(object sender, RoutedEventArgs e)
        {
            Application.Current.Shutdown();
        }

        private async void start_Click(object sender, RoutedEventArgs e)
        {
            // Закрытие всех процессов приложения
            Process[] processes = Process.GetProcessesByName("TimeManagement");
            foreach (Process process in processes)
            {
                process.Kill();
            }
            progressBar.Value = 0;

            //// Создаем экземпляр FileManager
            //FileManager fileManager = new FileManager();

            //// Создаем объект для отслеживания прогресса
            //IProgress<int> progress = new Progress<int>(p =>
            //{
            //    progressBar.Maximum = fileManager.CountOfFiles;
            //    // Обновляем значение прогресса в ProgressBar
            //    progressBar.Value = p;
            //});

            //// Вызываем асинхронный метод CompareAndDownloadFiles с прогрессом
            //await fileManager.CompareAndDownloadFiles(AppDomain.CurrentDomain.BaseDirectory, "", progress, UpdateDownloadStatuses);

            // Запуск приложения
            Process.Start("TimeManagement.exe");

            Application.Current.Shutdown();
        }

        private void UpdateDownloadStatuses(List<FileObserver> statuses)
        {
            // Обновление списка статусов загрузки
            DownloadStatuses.Clear();
            foreach (var status in statuses)
            {
                DownloadStatuses.Add(status);
            }
        }
    }
}