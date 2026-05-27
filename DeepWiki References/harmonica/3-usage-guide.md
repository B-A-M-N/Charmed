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

Usage Guide
===========

Relevant source files

*   [README.md](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1)
    
*   [harmonica.go](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go)
    

This guide provides practical instructions for integrating and using the Harmonica physics-based animation library in your applications. It covers the essential steps for implementing both spring animations and projectile motion simulations, with guidance on parameter selection and integration patterns.

For theoretical explanations of the underlying concepts, see [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts)
. For example implementations in different environments, see [Example Applications](https://deepwiki.com/charmbracelet/harmonica/4-example-applications)
.

Library Integration
-------------------

Harmonica is designed to be framework-agnostic, making it suitable for various Go applications including GUI, games, and terminal applications.

### Adding to Your Project

### Basic Integration Pattern

**Diagram: Harmonica Integration Pattern**

Sources: [harmonica.go4-32](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L4-L32)
 [README.md22-52](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L22-L52)

Spring Animation Usage
----------------------

Spring animations create smooth, natural motion by simulating a damped harmonic oscillator. This is useful for transitions, UI responses, and any animation that should feel natural.

### Initializing a Spring

**Diagram: Spring Initialization Process**

Sources: [README.md58-69](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L58-L69)
 [harmonica.go7-9](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L7-L9)

To create a spring:

1.  Determine your application's frame rate (for time delta)
2.  Choose appropriate angular frequency (speed) and damping ratio (springiness)
3.  Initialize using `NewSpring()`

Example:

### Spring Animation Loop

**Diagram: Spring Animation Loop**

Sources: [README.md31-49](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L31-L49)
 [harmonica.go10-16](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L10-L16)

For each frame in your animation loop:

1.  Call `spring.Update()` with current position, current velocity, and target position
2.  Use the returned new position and velocity values
3.  Render or update your UI with the new position

### Parameter Selection Guidelines

| Parameter | Typical Range | Effect | Common Use Cases |
| --- | --- | --- | --- |
| Time Delta | Based on FPS | Consistency of animation | Use `FPS(60)` for 60fps applications |
| Angular Frequency | 1.0 - 10.0 | Higher = faster animations | 4.0-8.0 for UI, 1.0-3.0 for slower animations |
| Damping Ratio | 0.1 - 2.0 | <1.0: oscillation, =1.0: no oscillation, fastest approach, >1.0: slower approach | 0.5-0.8 for slightly bouncy UI, 1.0 for critical damping |

Sources: [README.md58-91](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L58-L91)

### Spring Animation Example

Animate a sprite to a target position:

Sources: [README.md28-52](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L28-L52)

Projectile Motion Usage
-----------------------

Projectile motion simulates objects moving under constant acceleration, such as gravity. This is ideal for particle effects, ballistic animations, and physics simulations.

### Initializing a Projectile

**Diagram: Projectile Initialization Process**

Sources: [harmonica.go20-26](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L20-L26)

To create a projectile:

1.  Determine your application's frame rate (for time delta)
2.  Define initial position as a Point (x, y, z coordinates)
3.  Define initial velocity as a Vector (direction and magnitude)
4.  Define acceleration as a Vector (typically includes gravity)
5.  Initialize using `NewProjectile()`

Example:

### Projectile Animation Loop

**Diagram: Projectile Animation Loop**

Sources: [harmonica.go28-31](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L28-L31)

For each frame in your animation loop:

1.  Call `projectile.Update()` to advance the simulation
2.  Use the returned position to update your visual elements
3.  Optionally access current velocity with `projectile.Velocity()`

### Mapping Code to Visual Elements

**Diagram: Projectile Code to Visual Mapping**

Sources: [harmonica.go20-31](https://github.com/charmbracelet/harmonica/blob/6f6967c9/harmonica.go#L20-L31)

Performance Considerations
--------------------------

When using Harmonica in performance-sensitive applications:

1.  **Time Delta Accuracy**: Ensure your actual frame rate matches the time delta provided to Harmonica
2.  **Batch Updates**: For many objects using the same spring parameters, create one spring and use it to update multiple objects
3.  **Optimization Point**: Consider stopping the animation when position and velocity are close enough to the target to avoid unnecessary computation

**Diagram: Animation Optimization Pattern**

Sources: [README.md42-49](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L42-L49)

Common Use Cases
----------------

### Spring Animation

*   **UI Elements**: Smooth transitions for positioning, scaling, or opacity changes
*   **Camera Movement**: Natural-feeling camera tracking and transitions
*   **Input Feedback**: Responsive UI elements that react naturally to input

### Projectile Motion

*   **Particle Effects**: Realistic particles for explosions, water, etc.
*   **Object Trajectories**: Simulating thrown or launched objects
*   **Character Movement**: Jump arcs and movement paths for game characters

Troubleshooting
---------------

| Symptom | Possible Cause | Solution |
| --- | --- | --- |
| Animation too fast/slow | Incorrect angular frequency | Adjust angular frequency (higher = faster) |
| Too bouncy/rigid | Inappropriate damping ratio | Decrease damping ratio for more bounce, increase for less |
| Inconsistent animation | Time delta mismatch | Ensure time delta matches actual frame rate |
| Motion looks unnatural | Extreme parameter values | Use moderate parameter values, typically angular frequency: 2-8, damping: 0.2-1.2 |

Sources: [README.md58-91](https://github.com/charmbracelet/harmonica/blob/6f6967c9/README.md?plain=1#L58-L91)

For more detailed implementation examples, see [Spring Implementation](https://deepwiki.com/charmbracelet/harmonica/3.1-spring-implementation)
 and [Projectile Implementation](https://deepwiki.com/charmbracelet/harmonica/3.2-projectile-implementation)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#usage-guide)
    
*   [Library Integration](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#library-integration)
    
*   [Adding to Your Project](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#adding-to-your-project)
    
*   [Basic Integration Pattern](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#basic-integration-pattern)
    
*   [Spring Animation Usage](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#spring-animation-usage)
    
*   [Initializing a Spring](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#initializing-a-spring)
    
*   [Spring Animation Loop](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#spring-animation-loop)
    
*   [Parameter Selection Guidelines](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#parameter-selection-guidelines)
    
*   [Spring Animation Example](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#spring-animation-example)
    
*   [Projectile Motion Usage](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#projectile-motion-usage)
    
*   [Initializing a Projectile](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#initializing-a-projectile)
    
*   [Projectile Animation Loop](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#projectile-animation-loop)
    
*   [Mapping Code to Visual Elements](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#mapping-code-to-visual-elements)
    
*   [Performance Considerations](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#performance-considerations)
    
*   [Common Use Cases](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#common-use-cases)
    
*   [Spring Animation](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#spring-animation)
    
*   [Projectile Motion](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#projectile-motion)
    
*   [Troubleshooting](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide#troubleshooting)
