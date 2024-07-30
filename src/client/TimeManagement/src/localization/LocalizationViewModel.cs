using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TimeManagement.src.localization
{
    public class LocalizationViewModel : INotifyPropertyChanged
    {
        private Localizer _localization;
        private string _currentLanguage;
        private Dictionary<string, string> _currentTranslations;

        public event PropertyChangedEventHandler PropertyChanged;

        public string CurrentLanguage
        {
            get => _currentLanguage;
            set
            {
                if (_currentLanguage != value)
                {
                    _currentLanguage = value;
                    _currentTranslations = _localization.Translations[_currentLanguage];
                    OnPropertyChanged(string.Empty);
                }
            }
        }

        public string Login => _currentTranslations["login"];
        public string Password => _currentTranslations["password"];
        public string AuthButton => _currentTranslations["auth_button"];
        public string Cancel => _currentTranslations["cancel"];

        public LocalizationViewModel()
        {
            _localization = App.locale;
            CurrentLanguage = "en"; // Default
        }

        protected void OnPropertyChanged(string name)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
        }
    }

}
