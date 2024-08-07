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
        private static readonly HttpClient client = new HttpClient();
        private string url;

        public AuthModel(string url)
        {
            this.url = url + "/login";
        }
        public async Task<string> LoginAsync(string login, string password)
        {
            var data = new
            {
                Username = login,
                Password = password
            };

            var json = JsonConvert.SerializeObject(data);
            var content = new StringContent(json, Encoding.UTF8, "application/json");

            try
            {
                var response = await client.PostAsync(this.url, content);
                response.EnsureSuccessStatusCode();

                var responseBody = await response.Content.ReadAsStringAsync();
                return responseBody;
            }
            catch (HttpRequestException e)
            {
                var box = MessageBox.Error(e);
                await box.ShowAsync();
                return null;
            }
        }
    }
}
