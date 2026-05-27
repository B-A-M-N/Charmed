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

Core Concepts
=============

Relevant source files

*   [README.md](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1)
    
*   [harmonica.go](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go)
    

This document outlines the fundamental architecture and design principles behind the Harmonica library. Harmonica is a physics-based animation library that provides two primary systems: Spring simulations and Projectile motion. For specific implementation details, see [Spring Animation](https://deepwiki.com/charmbracelet/harmonica/2.1-spring-animation)
 and [Projectile Motion](https://deepwiki.com/charmbracelet/harmonica/2.3-projectile-motion)
.

Library Architecture
--------------------

The Harmonica library is designed with a simple, modular architecture that can be integrated into various application types, from terminal UIs to graphical 3D environments.

**Diagram: High-Level Architecture of Harmonica**

Sources: [harmonica.go1-32](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L1-L32)

Animation Systems Overview
--------------------------

The library contains two distinct animation systems, each designed for specific motion patterns:

| Animation System | Purpose | Key Parameters | Use Cases |
| --- | --- | --- | --- |
| **Spring** | Create smooth, natural spring-like motion | Time Delta, Angular Frequency, Damping Ratio | UI transitions, elastic motion |
| **Projectile** | Simulate objects moving under constant acceleration | Initial Position, Initial Velocity, Acceleration | Particles, ballistic motion, gravity |

Sources: [harmonica.go1-32](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L1-L32)
 [README.md11-25](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L11-L25)

Spring System
-------------

The Spring system implements a damped harmonic oscillator, providing realistic spring-like motion that can be tuned for various behaviors.

**Diagram: Spring System Structure**

### Spring Parameters

*   **Time Delta**: Time step for each animation update, often calculated using the `FPS()` utility
*   **Angular Frequency**: Controls the speed of the animation (higher values = faster motion)
*   **Damping Ratio**: Controls the "springiness" of the motion

Sources: [README.md57-69](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L57-L69)

### Spring Animation Flow

**Diagram: Spring Animation Data Flow**

The `Spring.Update()` method takes the current position, velocity, and target position, then calculates new position and velocity values based on the spring's parameters. These new values are typically used to update the animated object in each frame.

Sources: [README.md22-50](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L22-L50)
 [harmonica.go5-16](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L5-L16)

Projectile System
-----------------

The Projectile system simulates objects moving under constant acceleration (like gravity), allowing for realistic ballistic motion in 2D or 3D space.

**Diagram: Projectile System Structure**

### Projectile Animation Flow

**Diagram: Projectile Animation Data Flow**

The `Projectile.Update()` method advances the simulation by one time step, calculating a new position based on the current velocity and acceleration. The internal velocity is also updated but is typically managed by the Projectile object itself.

Sources: [harmonica.go18-31](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L18-L31)

Integration Pattern
-------------------

Harmonica is designed to be integrated into an application's animation loop, following a simple pattern:

**Diagram: Typical Integration Pattern**

This integration pattern applies to both Spring and Projectile animations, with minor variations depending on the specific system used.

Sources: [README.md22-52](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L22-L52)
 [harmonica.go5-31](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L5-L31)

Computational Model
-------------------

Harmonica uses physically accurate models for its animations:

1.  **Spring System**: Implements a damped harmonic oscillator based on [Ryan Juckett's approach](https://www.ryanjuckett.com/damped-springs/)
    , which precomputes coefficients for efficient updates.
    
2.  **Projectile System**: Uses standard Newtonian physics for projectile motion, calculating position changes based on velocity and acceleration vectors.
    

Both systems are designed to be computationally efficient, making them suitable for real-time applications with many animated objects.

Sources: [README.md93-100](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L93-L100)

Key Design Principles
---------------------

1.  **Framework Agnostic**: Harmonica doesn't depend on any specific rendering or UI framework.
2.  **Efficiency**: Precomputed coefficients and optimized algorithms ensure minimal computational overhead.
3.  **Simplicity**: Simple API with just a few methods makes integration straightforward.
4.  **Physically Accurate**: Based on proper physics models for realistic motion.
5.  **Flexible**: Works equally well in 2D, 3D, or terminal interfaces.

Sources: [README.md11-25](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L11-L25)
 [harmonica.go1-4](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L1-L4)

For detailed implementation guides on using these systems in your application, see [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#core-concepts)
    
*   [Library Architecture](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#library-architecture)
    
*   [Animation Systems Overview](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#animation-systems-overview)
    
*   [Spring System](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#spring-system)
    
*   [Spring Parameters](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#spring-parameters)
    
*   [Spring Animation Flow](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#spring-animation-flow)
    
*   [Projectile System](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#projectile-system)
    
*   [Projectile Animation Flow](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#projectile-animation-flow)
    
*   [Integration Pattern](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#integration-pattern)
    
*   [Computational Model](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#computational-model)
    
*   [Key Design Principles](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts#key-design-principles)
