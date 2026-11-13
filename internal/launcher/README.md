# Launcher implementation

This package implements the `xynaut` CLI and its Control API. The executable
and end-to-end command tests are in [`cmd/xynaut`](../../cmd/xynaut).

The launcher uses the public Semiont Go SDK as a dependency. The protocol
fixtures in this repository pin the contracts this version was extracted
against.
