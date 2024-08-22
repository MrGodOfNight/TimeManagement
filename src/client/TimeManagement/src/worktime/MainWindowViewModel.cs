using Avalonia.Controls;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Globalization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TimeManagement.src.localization;

namespace TimeManagement.src.worktime
{
    public class MainWindowViewModel : INotifyPropertyChanged
    {
        private Dictionary<string, string> _currentTranslations;

        public event PropertyChangedEventHandler PropertyChanged;

        public MainWindowViewModel(Dictionary<string, string> currentLang)
        {
            _currentTranslations = currentLang;
        }
        public string Main => _currentTranslations["main"];
        public string WorkTime => _currentTranslations["work_time"];
        public string BreakTime => _currentTranslations["break_time"];
        public string Break => _currentTranslations["break"];
        public string WorkButton => _currentTranslations["start_work"]; //_currentTranslations["end_work"];
        public string Report => _currentTranslations["report"];
        public string Send => _currentTranslations["send"];
        public string Admin => _currentTranslations["admin"];

        protected void OnPropertyChanged(string name)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
        }
    }
}
