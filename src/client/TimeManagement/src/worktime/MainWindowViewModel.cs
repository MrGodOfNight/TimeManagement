﻿/*
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
        private ObservableCollection<Table> _tables;
        public ObservableCollection<Table> Tables {
            get => _tables;
            set 
            {
                _tables = value;
                OnPropertyChanged(nameof(Tables));
            } 
        }
        private ObservableCollection<User> _users;
        public ObservableCollection<User> Users
        {
            get => _users;
            set
            {
                _users = value;
                OnPropertyChanged(nameof(Users));
            }
        }

        public event PropertyChangedEventHandler PropertyChanged;

        public MainWindowViewModel(Dictionary<string, string> currentLang)
        {
            _currentTranslations = currentLang;
        }
        public string Main => _currentTranslations["main"];
        public string WorkTime => _currentTranslations["work_time"];
        public string BreakTime => _currentTranslations["break_time"];
        public string StartBreak => _currentTranslations["start_break"];
        public string StopBreak => _currentTranslations["stop_break"];
        public string StartWork => _currentTranslations["start_work"];
        public string StopWork => _currentTranslations["end_work"];
        public string Report => _currentTranslations["report"];
        public string Send => _currentTranslations["send"];
        public string StatisticButton => _currentTranslations["statistic_button"];
        public string SendError => _currentTranslations["send_error"];
        public string Admin => _currentTranslations["admin"];
        public string Day => _currentTranslations["today"];
        public string Month => _currentTranslations["for_month"];
        public string Refresh => _currentTranslations["refresh"];
        public string Date1 => _currentTranslations["date"];
        public string Time1 => _currentTranslations["time"];
        public string LVL => _currentTranslations["adminlvl"];
        public string End => _currentTranslations["end"];
        public string Last => _currentTranslations["last_time"];
        public string Password1 => _currentTranslations["password"];
        public string Login1 => _currentTranslations["login"];
        protected void OnPropertyChanged(string name)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
        }
        public void UpdateTables(ObservableCollection<Table> tables)
        {
            this.Tables = tables;
        }
        public void UpdateUsers(ObservableCollection<User> users)
        {
            this.Users = users;
        }
    }
}
