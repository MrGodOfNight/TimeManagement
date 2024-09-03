using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TimeManagement.src.worktime
{
    public class Table
    {
        public string Date { get; set; }
        public string Time { get; set; }
        public override string  ToString()
        {
            return Date + Time;
        }
    }
    public class DayTime
    {
        public string Time { get; set; }
    }
}
