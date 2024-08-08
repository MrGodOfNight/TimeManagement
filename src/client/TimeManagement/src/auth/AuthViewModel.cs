/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

using MsBox.Avalonia.Enums;
using MsBox.Avalonia;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Avalonia.Controls;
using TimeManagement.src.localization;
using System.Globalization;

namespace TimeManagement.src.auth
{
    public class AuthViewModel : INotifyPropertyChanged
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
                    OnPropertyChanged(nameof(Auth));
                    OnPropertyChanged(nameof(WatermarkLogin));
                    OnPropertyChanged(nameof(WatermarkPassword));
                }
            }
        }

        public string Login => _currentTranslations["login"];
        public string Password => _currentTranslations["password"];
        public string AuthButton => _currentTranslations["auth_button"];
        public string Cancel => _currentTranslations["cancel"];
        public string Auth => _currentTranslations["auth"];
        public string WatermarkLogin => _currentTranslations["watermark_login"];
        public string WatermarkPassword => _currentTranslations["watermark_password"];
        public string Unauthorized => _currentTranslations["unauthorized"];
        public int SelectedLang { get; set; }

        public AuthViewModel()
        {
            string jsonContent = JsonManager.LoadJsonFile("TimeManagement.src.localization.localization.json");
            _localization = new Localizer(jsonContent);
            CultureInfo currentCulture = CultureInfo.CurrentCulture;
            switch (currentCulture.Name)
            {
                case "en-US":
                    _currentTranslations = _localization.Translations["en"];
                    SelectedLang = 0;
                    break;
                case "ru-RU":
                    _currentTranslations = _localization.Translations["ru"];
                    SelectedLang = 1;
                    break;
                default:
                    _currentTranslations = _localization.Translations["en"];
                    SelectedLang = 0;
                    break;
            }
        }

        protected void OnPropertyChanged(string name)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
        }
    }

}
