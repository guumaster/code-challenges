
## Developer Notes

----------------------------

These notes are for the reviewer to have more insight about how I coded this exercise and the design choices I made.

**Time commitment note**: I spent almost 7 hours to reach this solution, although I could have done a simpler version, I enjoyed doing it as it is and learn a few things along the way.


### Code structure & Design

This code base can be shared as a library that provides a decoupled and composable system to read and process logs.

Inside the `pkg` folder you will find: 

```
    pkg/
        monitor/    manager to control and coordinate the whole process.
        source/     input sources implemented and the Source interface.
        processor/  line processors and the Processor interfaces.

```

The `Monitor` object is a manager that can handle an input source (object implementing `Source` interface), and through this monitor each line will be sent to all registered processors (objects that implement `Processor` interface).

```
      Log Monitor App
      +-----------------------------------------------+
      |                  +---------+                  |
      |                  | Monitor |                  |
      |                  +------+--+                  |
      |                     ^   |    +-------------+  |
Input |  +--------+   Lines |   +--->| Processor 1 +--+--> Output Aggregation Stats
File -+->| Input  +---------+   |    +-------------+  |
      |  | Parser |             |                     |
      |  +--------+             |    +-------------+  |
      |                         +--->| Processor 2 +--+--> Output Alerting
      |                         |    +-------------+  |
      |                         |                     |
      |                         |    +-------------+  |
      |                         +--->| Processor N +--+--> Output N
      |                              +-------------+  |
      +-----------------------------------------------+
```

There are two main goal for this design, one is to separate each processor from the input source, and on the other hand, easy extensibility. As you can see in the `processor` folder you can create a new one with an easy interface and a lot of repetitive code already taken care of by the monitor. Same goes for the `source` folder where you can easily change from a csv parser to a database input by creating a new object and implement a simple interface.

There are two processor that are implemented but not requested `LinePrinter` and `GlobalStats`. The former is for debugging, and it prints all lines as they come through, and  the former to print the total line count at the end of the process. These two processor are only to show how to compose and pass it to the monitor.


### Tradeoffs and possible Improvements

- Most configuration parameters are fixed as required in the exercise' description. Located in  `cmd/log-monitor/fixed_params.go`, but could be opened as CLI params for more flexibility.
- The code assumes that all lines are sorted by timestamp, otherwise calculations may not be right.
- Logger utility is really simple and may be improved.
- The exercise hints to add debugging data on each 10 second logs, I'm not sure which problem is going to be debugged, hence, I wasn't sure what data to add.
- Due to time limitations I didn't add as many tests as I would do for a real project.
- The input file format is not dynamic, it has to comply with the example given, but with some work it could accept more line formats.
- A filtering mechanism could be implemented to analyze logs partially with arbitrary criteria (by time, by section, per IP, etc).
- As the solution is designed, changing it to parse log files in realtime is possible by adding a watcher and tailing every new line.
- Documentation on how to extend the CLI for new inputs and processors should be added to easy further development.
