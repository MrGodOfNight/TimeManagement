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

using Newtonsoft.Json.Linq;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace TimeManagement.src
{
    internal class RequestManager
    {
        private HttpClient _httpClient = new HttpClient();
        private string _token;
        public RequestManager(string token)
        {  _token = token; }
        public async Task<string> SendRequestPost(string uri, object data = null)
        {
            var json = JsonConvert.DeserializeObject<Dictionary<string, string>>(JsonManager.LoadJsonFile("TimeManagement.src.settings.json"));
            using var request = new HttpRequestMessage(HttpMethod.Post, json["server_uri"] + uri);
            request.Headers.Add("token", _token);
            if (data != null)
            {
                var json2 = JsonConvert.SerializeObject(data);
                request.Content = new StringContent(json2, Encoding.UTF8, "application/json");
            }
            try
            {
                using var response = await _httpClient.SendAsync(request);
                string responseData = await response.Content.ReadAsStringAsync();
                if (!response.IsSuccessStatusCode)
                {
                    var box = MessageBox.Error(responseData);
                    await box.ShowAsync();
                    return "";
                }
                return responseData;
            }
            catch (Exception ex)
            {

                var box = MessageBox.Error(ex);
                await box.ShowAsync();
                return "";
            }

        }
        public async Task<string> SendRequestGet(string uri, object data = null)
        {
            var json = JsonConvert.DeserializeObject<Dictionary<string, string>>(JsonManager.LoadJsonFile("TimeManagement.src.settings.json"));
            using var request = new HttpRequestMessage(HttpMethod.Get, json["server_uri"] + uri);
            request.Headers.Add("token", _token);
            if (data != null)
            {
                var json2 = JsonConvert.SerializeObject(data);
                request.Content = new StringContent(json2, Encoding.UTF8, "application/json");
            }
            try
            {
                using var response = await _httpClient.SendAsync(request);
                string responseData = await response.Content.ReadAsStringAsync();
                if (!response.IsSuccessStatusCode)
                {
                    var box = MessageBox.Error(responseData);
                    await box.ShowAsync();
                    return "";
                }
                return responseData;
            }
            catch (Exception ex)
            {

                var box = MessageBox.Error(ex);
                await box.ShowAsync();
                return "";
            }

        }
    }
}
