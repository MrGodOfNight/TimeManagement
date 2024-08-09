/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

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
                var response = await Program.client.PostAsync(this.url, content);
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
