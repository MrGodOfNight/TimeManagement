using Fizzler;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace TimeManagement.src.auth
{
    public class AuthModel
    {
        // Create the HttpClient instance
        private static readonly HttpClient client = new HttpClient();
        private string url;

        public AuthModel(string url)
        {
            // Set the URL for the login endpoint
            this.url = url + "/login";
        }
        public async Task<HttpResponseMessage> LoginAsync(string login, string password)
        {
            // Create a object to hold the data for the login request
            var data = new
            {
                Username = login,
                Password = password
            };

            // Serialize the data object to JSON
            var json = JsonConvert.SerializeObject(data);
            var content = new StringContent(json, Encoding.UTF8, "application/json");
            try
            {
                // Send the POST request to the login endpoint
                var response = await client.PostAsync(this.url, content);
                return response;
            }
            catch (Exception e)
            {
                var box = MessageBox.Error(e);
                await box.ShowAsync();
                return null;
            }

        }
    }
}
