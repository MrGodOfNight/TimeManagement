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
        // Dictionary to store the translations
        public Dictionary<string, Dictionary<string, string>> Translations { get; set; }
        public Localizer(string jsonContent)
        {
            // Deserialize the JSON content into a dictionary
            Translations = JsonConvert.DeserializeObject<Dictionary<string, Dictionary<string, string>>>(jsonContent);
        }
    }
}
