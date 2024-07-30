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

        public static Localizer DeserializeJson(string jsonContent)
        {
            return JsonConvert.DeserializeObject<Localizer>(jsonContent);
        }
        public static string LoadJsonFile(string resourceName)
        {
            var assembly = Assembly.GetExecutingAssembly();
            using (Stream stream = assembly.GetManifestResourceStream(resourceName))
            {
                if (stream == null)
                    throw new FileNotFoundException("Resource not found.", resourceName);

                using (StreamReader reader = new StreamReader(stream))
                {
                    return reader.ReadToEnd();
                }
            }
        }
    }
}
