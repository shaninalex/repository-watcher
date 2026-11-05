class Program
{
    static void Main(string[] args)
    {
        // 1) Read command line arguments and flags
        // it have to be arg - "--config" with the path to configuration file
        //  otherwise will be used ~/.config/repo-watcher/config.yaml - the default path.
        // 2) Read repositories list
        // 3) Check credentials ( ask for renew if needed )
        // 4) Start background API pooling tasks
        // 5) Open UDP for client
        // 6) Gathering events and send on client
    }
}