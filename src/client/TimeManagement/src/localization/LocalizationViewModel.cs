using MsBox.Avalonia.Enums;
using MsBox.Avalonia;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Avalonia.Controls;

namespace TimeManagement.src.localization
{
    public class LocalizationViewModel : INotifyPropertyChanged
    {
        private Localizer _localization;
        private ComboBoxItem _currentLanguage;
        private Dictionary<string, string> _currentTranslations;

        public event PropertyChangedEventHandler PropertyChanged;

        public ComboBoxItem CurrentLanguage
        {
            get => _currentLanguage;
            set
            {
                if (_currentLanguage != value)
                {
                    _currentLanguage = value;
                    _currentTranslations = _localization.Translations[_currentLanguage.Tag.ToString()];
                    OnPropertyChanged(nameof(_currentLanguage));
                    OnPropertyChanged(nameof(Login));
                    OnPropertyChanged(nameof(Password));
                    OnPropertyChanged(nameof(AuthButton));
                    OnPropertyChanged(nameof(Cancel));
                }
            }
        }

        public string Login => _currentTranslations["login"];
        public string Password => _currentTranslations["password"];
        public string AuthButton => _currentTranslations["auth_button"];
        public string Cancel => _currentTranslations["cancel"];

        public LocalizationViewModel()
        {
            string jsonContent = Localizer.LoadJsonFile("TimeManagement.src.localization.localization.json");
            _localization = new Localizer(jsonContent);
        }

        protected void OnPropertyChanged(string name)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
        }
    }

}
