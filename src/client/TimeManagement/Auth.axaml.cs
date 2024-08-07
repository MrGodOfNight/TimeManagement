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

namespace TimeManagement;

public partial class Auth : Window
{
    public Auth()
    {
        DataContext = new AuthViewModel();
        InitializeComponent();
    }
    public async void AuthHandler(object sender, RoutedEventArgs args)
    {
        var json = JsonConvert.DeserializeObject<Dictionary<string, string>>(JsonManager.LoadJsonFile("TimeManagement.src.settings.json"));
        AuthModel auth = new AuthModel(json["server_uri"]);
        var token = await auth.LoginAsync(UsernameTextBox.Text, PasswordTextBox.Text);
        var box = MessageBox.Debug(token);
        await box.ShowAsync();
    }
}