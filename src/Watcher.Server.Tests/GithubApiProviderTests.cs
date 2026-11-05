using System.Net;
using Watcher.Server.Core.providers.github;

namespace Watcher.Server.Tests;

public class GithubApiProviderTests
{
    private GithubApiProvider _client;

    [SetUp]
    public void Setup()
    {
        _client = new GithubApiProvider(null);
    }
    
    [Test]
    public async Task GetStatusOk_ListWorkflows()
    {
        HttpStatusCode response = await _client.ListWorkflows("shaninalex", "repository-watcher");
        Assert.That(response, Is.EqualTo(HttpStatusCode.OK));
    }

    [Test]
    public async Task GetStatusOk_ListRuns()
    {
        HttpStatusCode response = await _client.ListRuns("shaninalex", "repository-watcher");
        Assert.That(response, Is.EqualTo(HttpStatusCode.OK));
    }
}