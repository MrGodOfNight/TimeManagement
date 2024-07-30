using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TimeManagement.src.localization;

namespace TimeManagement
{
    public class AuthViewModel
    {
        public LocalizationViewModel LocalizationViewModel { get; }

        public AuthViewModel()
        {
            LocalizationViewModel = new LocalizationViewModel();
        }
    }

}
