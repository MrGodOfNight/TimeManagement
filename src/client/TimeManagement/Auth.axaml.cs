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

using Avalonia;
using Avalonia.Controls;
using Avalonia.Controls.Notifications;
using Avalonia.Markup.Xaml;
using MsBox.Avalonia.Enums;
using MsBox.Avalonia;
using TimeManagement.src.localization;
using System.Net.Http.Json;
using Avalonia.Interactivity;
using Newtonsoft.Json;
using System.Collections.Generic;
using System;
using TimeManagement.src.auth;
using TimeManagement.src;
using Newtonsoft.Json.Linq;
using System.Net.Http.Headers;
using System.Net.Http;
using TimeManagement.src.worktime;
using System.Collections.ObjectModel;

namespace TimeManagement;

public partial class Auth : Window
{
    public Auth()
    {
        // Define the data context for the window
        DataContext = new AuthViewModel();
        InitializeComponent();
    }
    public async void AuthHandler(object sender, RoutedEventArgs args)
    {
        if (DataContext is not AuthViewModel viewModel) return;
        AuthButton.IsEnabled = false;
        // Load the settings from the settings file
        var json = JsonConvert.DeserializeObject<Dictionary<string, string>>(JsonManager.LoadJsonFile("TimeManagement.src.settings.json"));
        // Create an instance of the AuthModel class
        AuthModel auth = new AuthModel(json["server_uri"]);
        // Login to the server with the provided username and password and get the token
        var response = await auth.LoginAsync(UsernameTextBox.Text, PasswordTextBox.Text);
        if (response == null)
        {
            AuthButton.IsEnabled = true;
            return;
        }
        var responseBody = await response.Content.ReadAsStringAsync();

        switch (response.StatusCode)
        {
            case System.Net.HttpStatusCode.Unauthorized:

                var box = MessageBox.Error(viewModel.Unauthorized);
                await box.ShowAsync();
                AuthButton.IsEnabled = true;
                return;
            default:
                try
                {
                    response.EnsureSuccessStatusCode();
                }
                catch (Exception)
                {
                    box = MessageBox.Error(responseBody);
                    await box.ShowAsync();
                    AuthButton.IsEnabled = true;
                    return;
                }
                break;
        }

        JObject jsonObject = JObject.Parse(responseBody);
        string token = jsonObject["token"].ToString();
        int admin = (int)jsonObject["admin"];
        RequestManager requestManager = new RequestManager(token);
        if (admin > 0)
        {
            var users = JsonConvert.DeserializeObject<ObservableCollection<User>>(await requestManager.SendRequestGet("/admin/getusers"));
            MainWindow mainWindow = new MainWindow(viewModel.CurrentTranslations, UsernameTextBox.Text, token, admin, users);
            mainWindow.Show();
        }
        else
        {
            MainWindow mainWindow = new MainWindow(viewModel.CurrentTranslations, UsernameTextBox.Text, token, admin);
            mainWindow.Show();
        }
        AuthButton.IsEnabled = true;
        this.Close();
    }
}