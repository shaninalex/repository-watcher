namespace Watcher.Shared.domain;

public enum RepositoryEventType
{
    ReleasePublished,
    WorkflowStarted,
    WorkflowCompleted,
    WorkflowFailed,
    PipelineStarted,
    PipelineCompleted,
    PipelineFailed
}

public sealed record RepositoryEvent(
    string Id,
    string Provider,
    string Repository,
    RepositoryEventType Type,
    string Title,
    string? Description,
    string? Url,
    string? Author,
    DateTimeOffset Timestamp
);