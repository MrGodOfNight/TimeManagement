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

namespace TimeManagement;

public partial class Auth : Window
{
    public Auth()
    {
        DataContext = new LocalizationViewModel();
        InitializeComponent();
    }
    public async void cl(object sender, RoutedEventArgs args)
    {
        var asd = new LocalizationViewModel();
        string jsonContent = Localizer.LoadJsonFile("TimeManagement.src.localization.localization.json");
        var _localization = new Localizer(jsonContent);
        var box = MessageBoxManager.GetMessageBoxStandard("Caption", _localization.Translations.ToString(), ButtonEnum.YesNo);
        await box.ShowAsync();
    }
}