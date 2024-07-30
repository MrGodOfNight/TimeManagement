using Avalonia;
using Avalonia.Controls;
using Avalonia.Markup.Xaml;
using TimeManagement.src.localization;

namespace TimeManagement;

public partial class Auth : Window
{
    public Auth()
    {
        InitializeComponent();
        DataContext = new AuthViewModel();
    }
}