using System.Net;
using System.Net.Http.Headers;

namespace Watcher.Server.Core.providers.github;

public class GithubApiProvider
{
    private HttpClient Client;

    public GithubApiProvider(string? githubToken)
    {
        Client = new HttpClient();
        Client.DefaultRequestHeaders.Add("Accept", "application/vnd.github+json");
        Client.DefaultRequestHeaders.Add("X-GitHub-Api-Version", "2026-03-10");
        
        string customUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36";
        Client.DefaultRequestHeaders.Add("User-Agent", customUserAgent);

        if (!string.IsNullOrEmpty(githubToken))
        {
            Client.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", githubToken);
        }
    }

    // list of workflows
    // https://api.github.com/repos/shaninalex/repository-watcher/actions/workflows
    // Docs: https://docs.github.com/en/rest/actions/workflows?apiVersion=2026-03-10
    public async Task<HttpStatusCode> ListWorkflows(string account, string repository)
    {
        string url = $"repos/{account}/{repository}/actions/workflows";
        return await ProcessRequest(url);
    }


    // All list of runs
    // https://api.github.com/repos/shaninalex/repository-watcher/actions/runs
    // Docs: https://docs.github.com/en/rest/actions/workflow-runs?apiVersion=2026-03-10#list-workflow-runs-for-a-repository
    public async Task<HttpStatusCode> ListRuns(string account, string repository)
    {
        string url = $"repos/{account}/{repository}/actions/runs";
        return await ProcessRequest(url);
    }

    private async Task<HttpStatusCode> ProcessRequest(string url)
    {
        try
        {
            Client.DefaultRequestHeaders.Accept.Clear();
            HttpResponseMessage response =
                await Client.GetAsync($"https://api.github.com/{url}");
            return response.StatusCode;
        }
        catch (HttpRequestException e)
        {
            Console.WriteLine("\nException Caught!");
            Console.WriteLine("Message :{0} ", e.Message);
            return HttpStatusCode.InternalServerError;
        }
    }
}