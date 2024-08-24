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

using Avalonia.Controls;
using Avalonia.Input;
using Avalonia.Interactivity;
using Avalonia.Threading;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using TimeManagement.src;
using TimeManagement.src.auth;
using TimeManagement.src.worktime;
using static System.Runtime.InteropServices.JavaScript.JSType;

namespace TimeManagement
{
    public partial class MainWindow : Window
    {
        private HttpClient _httpClient = new HttpClient();
        private string _token;
        private MainWindowViewModel _viewModel;
        private DispatcherTimer _workTimer;
        private DispatcherTimer _breakTimer;
        private TimeSpan _workTime;
        private TimeSpan _breakTime;
        private int _workClickCount = 0;
        private int _breakClickCount = 0;
        private int _workTimeID;
        private int _breakTimeID;
        private int _adminLvl;
        public MainWindow(Dictionary<string, string> currentLang, string token, int adminLvl)
        {
            _token = token;
            _adminLvl = adminLvl;
            _viewModel = new MainWindowViewModel(currentLang);
            DataContext = _viewModel;
            InitializeComponent();
            if (_adminLvl < 1)
            {
                AdminPanel.IsEnabled = false;
            }
            Report.IsEnabled = false;
            SendButton.IsEnabled = false;
            BreakButton.IsEnabled = false;

            // init timer for work time
            _workTimer = new DispatcherTimer { Interval = TimeSpan.FromSeconds(1) };
            _workTimer.Tick += WorkTimer_Tick;

            // init timer for break time
            _breakTimer = new DispatcherTimer { Interval = TimeSpan.FromSeconds(1) };
            _breakTimer.Tick += BreakTimer_Tick;
        }
        private void WorkTimer_Tick(object sender, EventArgs e)
        {
            _workTime = _workTime.Add(TimeSpan.FromSeconds(1));
            WorkTimer.Text = _workTime.ToString(@"hh\:mm\:ss");
        }
        private void BreakTimer_Tick(object sender, EventArgs e)
        {
            _breakTime = _breakTime.Add(TimeSpan.FromSeconds(1));
            BreakTimer.Text = _breakTime.ToString(@"hh\:mm\:ss");
        }
        public async void WorkTimer_Click(object sender, RoutedEventArgs e)
        {
            _workClickCount++;
            
            if (_workClickCount == 1)
            {
                var res = await SendRequest("/work/start", new
                {
                    Time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"),
                });
                if (res == "" && res != null) return;
                else
                {
                    JObject jsonObject = JObject.Parse(res);
                    _workTimeID = (int)jsonObject["time_id"]; 
                }
                _breakTimer.Stop();
                _workTimer.Start();
                Report.IsEnabled = true;
                SendButton.IsEnabled = true;
                BreakButton.IsEnabled = true;
                WorkButton.Content = _viewModel.StopWork;
            }
            else if (_workClickCount == 2)
            {
                if(_breakClickCount == 1)
                {
                    if (await SendRequest("/break/stop", new
                    {
                        Time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"),
                        ID = _breakTimeID,
                    }) == "") return;
                }
                if (await SendRequest("/work/stop", new
                {
                    Time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"),
                    ID = _workTimeID,
                }) == "") return;
                _workTimer.Stop();
                _breakTimer.Stop();
                _workTime = TimeSpan.Zero;
                _breakTime = TimeSpan.Zero;
                WorkButton.Content = _viewModel.StartWork;
                Report.IsEnabled = false;
                SendButton.IsEnabled = false;
                BreakButton.IsEnabled = false;
                _workClickCount = 0;
            }
        }
        public async void BreakTimer_Click(object sender, RoutedEventArgs e)
        {
            _breakClickCount++;

            if (_breakClickCount == 1)
            {
                var res = await SendRequest("/break/start", new
                {
                    Time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"),
                });
                if (res == "" && res != null) return;
                else
                {
                    JObject jsonObject = JObject.Parse(res);
                    _breakTimeID = (int)jsonObject["time_id"];
                }
                _workTimer.Stop();
                _breakTimer.Start();
                BreakButton.Content = _viewModel.StopBreak;
            }
            else if (_breakClickCount == 2)
            {
                if (await SendRequest("/break/stop", new
                {
                    Time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"),
                    ID = _breakTimeID,
                }) == "") return;
                _breakTimer.Stop();
                _workTimer.Start();
                BreakButton.Content = _viewModel.StartBreak;
                _breakClickCount = 0;
            }
        }
        public async void Send(object sender, RoutedEventArgs e)
        {
            if (Report.Text == "")
            {
                var box = MessageBox.Error(_viewModel.SendError);
                await box.ShowAsync();
            }
            else
            {
                if (await SendRequest("/work/report", new
                {
                    ID = _workTimeID,
                    Report.Text,
                }) == "") return;
            }
        }

        private async Task<string> SendRequest(string uri, object data)
        {
            var json = JsonConvert.DeserializeObject<Dictionary<string, string>>(JsonManager.LoadJsonFile("TimeManagement.src.settings.json"));
            using var request = new HttpRequestMessage(HttpMethod.Post, json["server_uri"] + uri);
            request.Headers.Add("token", _token);
            var json2 = JsonConvert.SerializeObject(data);
            request.Content = new StringContent(json2, Encoding.UTF8, "application/json");
            using var response = await _httpClient.SendAsync(request);
            string responseData = await response.Content.ReadAsStringAsync();
            if (!response.IsSuccessStatusCode)
            {
                var box = MessageBox.Error(responseData);
                await box.ShowAsync();
                return "";
            }
            return responseData;
        }
    }
}