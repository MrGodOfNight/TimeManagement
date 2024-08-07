using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;

namespace TimeManagement.src.localization
{
    public class Localizer
    {
        public Dictionary<string, Dictionary<string, string>> Translations { get; set; }
        public Localizer(string jsonContent)
        {
            Translations = JsonConvert.DeserializeObject<Dictionary<string, Dictionary<string, string>>>(jsonContent);
        }
    }
}
