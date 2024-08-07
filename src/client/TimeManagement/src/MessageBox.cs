using Avalonia.Controls;
using MsBox.Avalonia.Dto;
using MsBox.Avalonia.Models;
using MsBox.Avalonia;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;
using MsBox.Avalonia.Base;

namespace TimeManagement.src
{
    public class MessageBox
    {
        public static IMsBox<string> Debug(string text)
        {
            return MessageBoxManager.GetMessageBoxCustom(
            new MessageBoxCustomParams
            {
                ButtonDefinitions = new List<ButtonDefinition>
                {
                    new ButtonDefinition { Name = "Ok", }
                },
                ContentTitle = "Debug",
                ContentMessage = text,
                Icon = MsBox.Avalonia.Enums.Icon.Setting,
                WindowStartupLocation = WindowStartupLocation.CenterOwner,
                CanResize = true,
                MaxWidth = 600,
                MaxHeight = 800,
                SizeToContent = SizeToContent.WidthAndHeight,
                ShowInCenter = true,
                Topmost = false,
            });
        }
        public static IMsBox<string> Warning(string text)
        {
            return MessageBoxManager.GetMessageBoxCustom(
            new MessageBoxCustomParams
            {
                ButtonDefinitions = new List<ButtonDefinition>
                {
                    new ButtonDefinition { Name = "Ok", }
                },
                ContentTitle = "Warning",
                ContentMessage = text,
                Icon = MsBox.Avalonia.Enums.Icon.Warning,
                WindowStartupLocation = WindowStartupLocation.CenterOwner,
                CanResize = true,
                MaxWidth = 600,
                MaxHeight = 800,
                SizeToContent = SizeToContent.WidthAndHeight,
                ShowInCenter = true,
                Topmost = false,
            });
        }
        public static IMsBox<string> Error(Exception e)
        {
            return MessageBoxManager.GetMessageBoxCustom(
            new MessageBoxCustomParams
            {
                ButtonDefinitions = new List<ButtonDefinition>
                {
                    new ButtonDefinition { Name = "Ok", }
                },
                ContentTitle = "Error",
                ContentMessage = e.Message,
                Icon = MsBox.Avalonia.Enums.Icon.Error,
                WindowStartupLocation = WindowStartupLocation.CenterOwner,
                CanResize = true,
                MaxWidth = 600,
                MaxHeight = 800,
                SizeToContent = SizeToContent.WidthAndHeight,
                ShowInCenter = true,
                Topmost = false,
            });
        }
    }
}
