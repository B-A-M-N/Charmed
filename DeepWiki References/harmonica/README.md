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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1)
    

Harmonica is a physics-based animation library in Go that provides smooth, natural motion through spring simulations and projectile motion. This document introduces the library's purpose, architecture, and core components.

For detailed information about implementing spring animations, see [Spring Implementation](https://deepwiki.com/charmbracelet/harmonica/3.1-spring-implementation)
. For projectile motion implementation, see [Projectile Implementation](https://deepwiki.com/charmbracelet/harmonica/3.2-projectile-implementation)
.

Purpose and Scope
-----------------

Harmonica enables developers to create animations that follow physical principles, resulting in motion that feels organic and satisfying. The library is:

*   **Framework-agnostic**: Works well in 2D, 3D, and terminal-based interfaces
*   **Efficient**: Uses precomputed coefficients for optimal performance
*   **Simple**: Provides a straightforward API requiring minimal setup

The library focuses on two primary physics-based systems:

1.  **Spring Animation**: For smooth transitions following damped harmonic motion
2.  **Projectile Motion**: For simulating objects moving under constant acceleration

Sources: [README.md11-17](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L11-L17)
 [README.md24-27](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L24-L27)

Architecture Overview
---------------------

### High-Level Structure

Harmonica's architecture is organized around its two main simulation systems:

Sources: [README.md22-55](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L22-L55)

### Core Components Relationship

The library provides specialized simulation components for different types of physical motion:

Sources: [README.md25-27](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L25-L27)
 [README.md42-47](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L42-L47)

Spring System
-------------

The Spring system simulates damped harmonic oscillation - the core feature of Harmonica.

### Spring Configuration

The Spring is initialized with three key parameters:

| Parameter | Description | Typical Values |
| --- | --- | --- |
| Delta Time | Time step for the simulation | Use `FPS(60)` for 60fps |
| Angular Frequency | Controls animation speed | 2.0-10.0 (higher is faster) |
| Damping Ratio | Controls "springiness" | 0.0-1.0+ (lower is springier) |

Sources: [README.md59-69](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L59-L69)

### Damping Behavior

The damping ratio controls how the spring approaches equilibrium:

| Damping Type | Ratio Value | Behavior |
| --- | --- | --- |
| Under-damped | < 1.0 | Oscillates before reaching equilibrium (springy) |
| Critically-damped | \= 1.0 | Fastest approach without oscillation |
| Over-damped | \> 1.0 | Gradual approach without oscillation |

For more details about damping types, see [Damping Types](https://deepwiki.com/charmbracelet/harmonica/2.2-damping-types)
.

Sources: [README.md71-91](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L71-L91)

Projectile Motion
-----------------

The Projectile system simulates objects moving under constant acceleration (like gravity) in 2D or 3D space.

### Projectile Components

A projectile tracks three key properties:

*   **Position**: Current location in space (X, Y, Z)
*   **Velocity**: Speed and direction of movement
*   **Acceleration**: Rate of change in velocity (often gravity)

The system is ideal for:

*   Particle effects
*   Object trajectories
*   Physics-based simulations

Animation Flow
--------------

### Spring Animation Data Flow

On each frame, the application calls `Spring.Update()` with the current position, velocity, and target position, then receives the new position and velocity to apply.

Sources: [README.md44-48](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L44-L48)

### Projectile Motion Data Flow

Integration with Applications
-----------------------------

Harmonica is designed to be integrated into an application's animation loop:

Example integration code for spring animation:

The library works with various rendering systems:

| Application Type | Example |
| --- | --- |
| Terminal UI | Using Bubble Tea ([Terminal UI Example](https://deepwiki.com/charmbracelet/harmonica/4.1-terminal-ui-example)<br>) |
| 2D Graphics | Using OpenGL/Pixel ([OpenGL Example](https://deepwiki.com/charmbracelet/harmonica/4.2-opengl-example)<br>) |
| Particle Systems | For projectile motion ([Particle Example](https://deepwiki.com/charmbracelet/harmonica/4.3-particle-example)<br>) |

Sources: [README.md28-50](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L28-L50)

Summary
-------

Harmonica provides physics-based animation capabilities through its Spring and Projectile systems. By simulating natural physical processes, it creates animations that look and feel organic rather than mechanical.

The spring animation system is based on Ryan Juckett's damped simple harmonic oscillator implementation, adapted to Go for easy integration with various applications.

For deeper understanding of the core concepts, see [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts)
. For practical implementation guidance, see the [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
.

Sources: [README.md94-100](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L94-L100)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/harmonica#overview)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/harmonica#purpose-and-scope)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/harmonica#architecture-overview)
    
*   [High-Level Structure](https://deepwiki.com/charmbracelet/harmonica#high-level-structure)
    
*   [Core Components Relationship](https://deepwiki.com/charmbracelet/harmonica#core-components-relationship)
    
*   [Spring System](https://deepwiki.com/charmbracelet/harmonica#spring-system)
    
*   [Spring Configuration](https://deepwiki.com/charmbracelet/harmonica#spring-configuration)
    
*   [Damping Behavior](https://deepwiki.com/charmbracelet/harmonica#damping-behavior)
    
*   [Projectile Motion](https://deepwiki.com/charmbracelet/harmonica#projectile-motion)
    
*   [Projectile Components](https://deepwiki.com/charmbracelet/harmonica#projectile-components)
    
*   [Animation Flow](https://deepwiki.com/charmbracelet/harmonica#animation-flow)
    
*   [Spring Animation Data Flow](https://deepwiki.com/charmbracelet/harmonica#spring-animation-data-flow)
    
*   [Projectile Motion Data Flow](https://deepwiki.com/charmbracelet/harmonica#projectile-motion-data-flow)
    
*   [Integration with Applications](https://deepwiki.com/charmbracelet/harmonica#integration-with-applications)
    
*   [Summary](https://deepwiki.com/charmbracelet/harmonica#summary)
