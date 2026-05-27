Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/harmonica](https://github.com/charmbracelet/harmonica "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 24 April 2025 ([6f6967](https://github.com/charmbracelet/harmonica/commits/6f6967c9)
)

*   [Overview](https://deepwiki.com/charmbracelet/harmonica/1-overview)
    
*   [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts)
    
*   [Spring Animation](https://deepwiki.com/charmbracelet/harmonica/2.1-spring-animation)
    
*   [Damping Types](https://deepwiki.com/charmbracelet/harmonica/2.2-damping-types)
    
*   [Projectile Motion](https://deepwiki.com/charmbracelet/harmonica/2.3-projectile-motion)
    
*   [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
    
*   [Spring Implementation](https://deepwiki.com/charmbracelet/harmonica/3.1-spring-implementation)
    
*   [Projectile Implementation](https://deepwiki.com/charmbracelet/harmonica/3.2-projectile-implementation)
    
*   [Example Applications](https://deepwiki.com/charmbracelet/harmonica/4-example-applications)
    
*   [Terminal UI Example](https://deepwiki.com/charmbracelet/harmonica/4.1-terminal-ui-example)
    
*   [OpenGL Example](https://deepwiki.com/charmbracelet/harmonica/4.2-opengl-example)
    
*   [Particle Example](https://deepwiki.com/charmbracelet/harmonica/4.3-particle-example)
    
*   [Contributing](https://deepwiki.com/charmbracelet/harmonica/5-contributing)
    

Menu

Example Applications
====================

Relevant source files

*   [examples/go.mod](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod)
    
*   [examples/go.sum](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.sum)
    

Purpose and Scope
-----------------

This page provides an overview of the example applications included with the Harmonica library. These examples demonstrate how to integrate Harmonica's physics-based animation capabilities into various environments and frameworks. They serve as practical demonstrations of the concepts explained in the [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts)
 and [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
 sections.

Overview of Example Applications
--------------------------------

Harmonica includes three primary example applications that showcase different use cases for the library:

Sources: [examples/go.mod7-14](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L7-L14)

### Terminal UI Example

The Terminal UI example demonstrates using spring animations in a text-based user interface. It leverages the Bubble Tea library to create an interactive terminal application with smooth, spring-based transitions.

Sources: [examples/go.mod8-10](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L8-L10)

### OpenGL Example

The OpenGL example showcases spring animation in a graphical context, using the Pixel library (which is built on OpenGL) to render smooth visual transitions powered by Harmonica's spring physics.

Sources: [examples/go.mod11-12](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L11-L12)

### Particle Example

The Particle example demonstrates projectile motion simulation, showing how to use Harmonica to create realistic particle systems with gravity and other forces.

Sources: [examples/go.mod9](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L9-L9)

Integration Patterns
--------------------

Each example application demonstrates a different pattern for integrating Harmonica into an application's animation loop. The following diagram illustrates these integration patterns:

Sources: [examples/go.mod7-14](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L7-L14)

Example Application Dependencies
--------------------------------

The example applications rely on several external libraries to provide rendering capabilities and framework support:

| Example | Primary Dependencies | Purpose |
| --- | --- | --- |
| Terminal UI | `github.com/charmbracelet/bubbletea` | Terminal UI framework |
| Terminal UI | `github.com/charmbracelet/lipgloss` | Terminal styling |
| OpenGL | `github.com/faiface/pixel` | 2D graphics library |
| OpenGL | `github.com/fogleman/gg` | Graphics rendering |
| All Examples | `github.com/charmbracelet/harmonica` | Physics-based animation |

Sources: [examples/go.mod7-14](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L7-L14)

Code Structure Relationship
---------------------------

This diagram shows how the example applications relate to the core Harmonica code structure:

Sources: [examples/go.mod7-14](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L7-L14)

Running the Examples
--------------------

The examples are designed to demonstrate different aspects of the Harmonica library in various environments. Each example can be run from its respective directory:

Sources: [examples/go.mod1-14](https://github.com/charmbracelet/harmonica/blob/6f6967c9/examples/go.mod#L1-L14)

Summary
-------

The example applications included with Harmonica demonstrate how to integrate the library's physics-based animation capabilities into different environments:

1.  **Terminal UI Example**: Shows how to use spring animations in a text-based interface using Bubble Tea
2.  **OpenGL Example**: Demonstrates spring animations in a graphical context using Pixel/OpenGL
3.  **Particle Example**: Showcases projectile motion simulation for particle effects

These examples serve as practical implementations of the concepts covered in the [Spring Implementation](https://deepwiki.com/charmbracelet/harmonica/3.1-spring-implementation)
 and [Projectile Implementation](https://deepwiki.com/charmbracelet/harmonica/3.2-projectile-implementation)
 guides, providing working code that can be used as a reference when integrating Harmonica into your own applications.

For more detailed explanations of each example, see:

*   [Terminal UI Example](https://deepwiki.com/charmbracelet/harmonica/4.1-terminal-ui-example)
    
*   [OpenGL Example](https://deepwiki.com/charmbracelet/harmonica/4.2-opengl-example)
    
*   [Particle Example](https://deepwiki.com/charmbracelet/harmonica/4.3-particle-example)
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Example Applications](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#example-applications)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#purpose-and-scope)
    
*   [Overview of Example Applications](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#overview-of-example-applications)
    
*   [Terminal UI Example](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#terminal-ui-example)
    
*   [OpenGL Example](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#opengl-example)
    
*   [Particle Example](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#particle-example)
    
*   [Integration Patterns](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#integration-patterns)
    
*   [Example Application Dependencies](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#example-application-dependencies)
    
*   [Code Structure Relationship](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#code-structure-relationship)
    
*   [Running the Examples](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#running-the-examples)
    
*   [Summary](https://deepwiki.com/charmbracelet/harmonica/4-example-applications#summary)
