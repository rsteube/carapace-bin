package action

import "github.com/rsteube/carapace"

func ActionAdvanced() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionAdvancedKeys()
		case 1:
			return ActionAdvancedValues(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionAdvancedKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"+DisableAttachMechanism", "Enables the option that disables the mechanism that lets tools attach to the JVM",
		"ErrorFile=", "Specifies the path and file name to which error data is written when an irrecoverable error occurs",
		"+FailOverToOldVerifier", "Enables automatic failover to the old verifier when the new type checker fails",
		"LargePageSizeInBytes=", "On Solaris, sets the maximum size for large pages used for Java heap",
		"MaxDirectMemorySize=", "Sets the maximum total size direct-buffer allocations",
		"NativeMemoryTracking=", "Specifies the mode for tracking JVM native memory usage",
		"ObjectAlignmentInBytes=", "Sets the memory alignment of Java objects",
		"OnError=", "Sets a custom command or a series of semicolon-separated commands to run when an irrecoverable error occurs",
		"OnOutOfMemoryError=", "Sets a custom command or a series of semicolon-separated commands to run when an OutOfMemoryError exception is first thrown",
		"+PerfDataSaveToFile", "If enabled, saves jstat binary data when the Java application exits",
		"+PrintCommandLineFlags", "Enables printing of ergonomically selected JVM flags that appeared on the command line",
		"+PrintNMTStatistics", "Enables printing of collected native memory tracking data at JVM exit when native memory tracking is enabled",
		"+RelaxAccessControlCheck", "Decreases the amount of access control checks in the verifier",
		"+ShowMessageBoxOnError", "Enables displaying of a dialog box when the JVM experiences an irrecoverable error",
		"ThreadStackSize=", "Sets the thread stack size",
		"+TraceClassLoading", "Enables tracing of classes as they are loaded",
		"+TraceClassLoadingPreorder", "Enables tracing of all loaded classes in the order in which they are referenced",
		"+TraceClassResolution", "Enables tracing of constant pool resolutions",
		"+TraceClassUnloading", "Enables tracing of classes as they are unloaded",
		"+TraceLoaderConstraints", "Enables tracing of the loader constraints recording",
		"+UseAltSigs", "Enables the use of alternative signals instead of SIGUSR1 and SIGUSR2 for JVM internal signals",
		"-UseBiasedLocking", "Disables the use of biased locking",
		"-UseCompressedOops", "Disables the use of compressed pointers",
		"+UseHugeTLBFS", "This option pre-allocates all large pages up-front",
		"+UseLargePages", "Enables the use of large page memory",
		"+UseMembar", "Enables issuing of membars on thread state transitions",
		"+UsePerfData", "Enables the perfdata feature",
		"+UseTransparentHugePages", "On Linux, enables the use of large pages that can dynamically grow or shrink",
		"+AllowUserSignalHandlers", "Enables installation of signal handlers by the application",
		"+AggressiveOpts", "Enables the use of aggressive performance optimization features",
		"AllocateInstancePrefetchLines=", "Sets the number of lines to prefetch ahead of the instance allocation pointer",
		"AllocatePrefetchDistance=", "Sets the size, of the prefetch distance for object allocation",
		"AllocatePrefetchInstr=", "Sets the prefetch instruction to prefetch ahead of the allocation pointer",
		"AllocatePrefetchLines=", "Sets the number of cache lines to load after the last object allocation by using the prefetch instructions generated in compiled code",
		"AllocatePrefetchStepSize=", "Sets the step size for sequential prefetch instructions",
		"AllocatePrefetchStyle=", "Sets the generated code style for prefetch instructions",
		"+BackgroundCompilation", "Enables background compilation",
		"CICompilerCount=", "Sets the number of compiler threads to use for compilation",
		"CodeCacheMinimumFreeSpace=", "Sets the minimum free space required for compilation",
		"CompileCommand=", "Specifies a command to perform on a method",
		"CompileCommandFile=", "Sets the file from which JIT compiler commands are read",
		"CompileOnly=", "Sets the list of methods to which compilation should be restricted",
		"CompileThreshold=", "Sets the number of interpreted method invocations before compilation",
		"+DoEscapeAnalysis", "Enables the use of escape analysis",
		"InitialCodeCacheSize=", "Sets the initial code cache size",
		"+Inline", "Enables method inlining",
		"InlineSmallCode=", "Sets the maximum code size for compiled methods that should be inlined",
		"+LogCompilation", "Enables logging of compilation activity to a file named hotspot.log in the current working directory",
		"MaxInlineSize=", "Sets the maximum bytecode size  of a method to be inlined",
		"MaxNodeLimit=", "Sets the maximum number of nodes to be used during single method compilation",
		"MaxTrivialSize=", "Sets the maximum bytecode size of a trivial method to be inlined",
		"+OptimizeStringConcat", "Enables the optimization of String concatenation operations",
		"+PrintAssembly", "Enables printing of assembly code for bytecoded and native methods by using the external disassembler.so library",
		"+PrintCompilation", "Enables verbose diagnostic output from the JVM by printing a message to the console every time a method is compiled",
		"+PrintInlining", "Enables printing of inlining decisions",
		"ReservedCodeCacheSize=", "Sets the maximum code cache size  for JIT-compiled code",
		"RTMAbortRatio=", "The RTM abort ratio is specified as a percentage of all executed RTM transactions",
		"RTMRetryCount=", "RTM locking code will be retried, when it is aborted or busy, the number of times specified by this option",
		"-TieredCompilation", "Disables the use of tiered compilation",
		"+UseAES", "Enables hardware-based AES intrinsics for Intel, AMD, and SPARC hardware",
		"+UseAESIntrinsics", "UseAES and UseAESIntrinsics flags are enabled by default and are supported only for Java HotSpot Server VM 32-bit and 64-bit",
		"+UseCodeCacheFlushing", "Enables flushing of the code cache before shutting down the compiler",
		"+UseCondCardMark", "Enables checking of whether the card is already marked before updating the card table",
		"+UseRTMDeopt", "Auto-tunes RTM locking depending on the abort ratio",
		"+UseRTMLocking", "Generate Restricted Transactional Memory locking code for all inflated locks",
		"+UseSHA", "Enables hardware-based intrinsics for SHA crypto hash functions for SPARC hardware",
		"+UseSHA1Intrinsics", "Enables intrinsics for SHA-1 crypto hash function",
		"+UseSHA256Intrinsics", "Enables intrinsics for SHA-224 and SHA-256 crypto hash functions",
		"+UseSHA512Intrinsics", "Enables intrinsics for SHA-384 and SHA-512 crypto hash functions",
		"+UseSuperWord", "Enables the transformation of scalar operations into superword operations",
		"+ExtendedDTraceProbes", "Enables additional dtrace tool probes that impact the performance",
		"+HeapDumpOnOutOfMemory", "Enables the dumping of the Java heap to a file in the current directory",
		"HeapDumpPath=", "Sets the path and file name for writing the heap dump provided by the heap profiler",
		"LogFile=", "Sets the path and file name where log data is written",
		"+PrintClassHistogram", "Enables printing of a class instance histogram after a Control+C event",
		"+PrintConcurrentLocks", "Enables printing of locks after a event",
		"+UnlockDiagnosticVMOptions", "Unlocks the options intended for diagnosing the JVM",
		"nced Garbage Collection Options", "These options control how garbage collection is performed by the Java HotSpot VM.",
		"+AggressiveHeap", "Enables Java heap optimization",
		"+AlwaysPreTouch", "Enables touching of every page on the Java heap during JVM initialization",
		"+CMSClassUnloadingEnabled", "Enables class unloading when using the concurrent mark-sweep garbage collector",
		"CMSExpAvgFactor=", "Sets the percentage of time used to weight the current sample",
		"CMSInitiatingOccupancyFraction=", "Sets the percentage of the old generation occupancy at which to start a CMS collection cycle",
		"+CMSScavengeBeforeRemark", "Enables scavenging attempts before the CMS remark step",
		"CMSTriggerRatio=", "Sets the percentage of the value specified by -XX:MinHeapFreeRatio that is allocated before a CMS collection cycle commences",
		"ConcGCThreads=", "Sets the number of threads used for concurrent GC",
		"+DisableExplicitGC", "Enables the option that disables processing of calls to System.gc()",
		"+ExplicitGCInvokesConcurrent", "Enables invoking of concurrent GC by using the System.gc() request",
		"+ExplicitGCInvokesConcurrentAndUnloadsClasses", "Enables invoking of concurrent GC",
		"G1HeapRegionSize=", "Sets the size of the regions into which the Java heap is subdivided",
		"+G1PrintHeapRegions", "Enables the printing of information about which regions are allocated",
		"G1ReservePercent=", "Sets the percentage of the heap that is reserved",
		"InitialHeapSize=", "Sets the initial size of the memory allocation pool",
		"InitialSurvivorRatio=", "Sets the initial survivor space ratio used by the throughput garbage collector",
		"InitiatingHeapOccupancyPercent=", "Sets the percentage of the heap occupancy",
		"MaxGCPauseMillis=", "Sets a target for the maximum GC pause time",
		"MaxHeapSize=", "Sets the maximum size of the memory allocation pool",
		"MaxHeapFreeRatio=", "Sets the maximum allowed percentage of free heap space",
		"MaxMetaspaceSize=", "Sets the maximum amount of native memory that can be allocated for class metadata",
		"MaxNewSize=", "Sets the maximum size of the heap for the young generation",
		"MaxTenuringThreshold=", "Sets the maximum tenuring threshold for use in adaptive GC sizing",
		"MetaspaceSize=", "Sets the size of the allocated class metadata space that will trigger a garbage collection",
		"MinHeapFreeRatio=", "Sets the minimum allowed percentage of free heap space after a GC event",
		"NewRatio=", "Sets the ratio between young and old generation sizes",
		"NewSize=", "Sets the initial size of the heap for the young generation",
		"ParallelGCThreads=", "Sets the number of threads used for parallel garbage collection in the young and old generations",
		"+ParallelRefProcEnabled", "Enables parallel reference processing",
		"+PrintAdaptiveSizePolicy", "Enables printing of information about adaptive generation sizing",
		"+PrintGC", "Enables printing of messages at every GC",
		"+PrintGCApplicationConcurrentTime", "Enables printing of how much time elapsed since the last pause",
		"+PrintGCApplicationStoppedTime", "Enables printing of how much time the pause lasted",
		"+PrintGCDateStamps", "Enables printing of a date stamp at every GC",
		"+PrintGCDetails", "Enables printing of detailed messages at every GC",
		"+PrintGCTaskTimeStamps", "Enables printing of time stamps for every individual GC worker thread task",
		"+PrintGCTimeStamps", "Enables printing of time stamps at every GC",
		"+PrintStringDeduplicationStatistics", "Prints detailed deduplication statistics",
		"+PrintTenuringDistribution", "Enables printing of tenuring age information",
		"+ScavengeBeforeFullGC", "Enables GC of the young generation before each full GC",
		"SoftRefLRUPolicyMSPerMB=", "Sets the amount of time a softly reachable object is kept active on the heap",
		"StringDeduplicationAgeThreshold=", "String objects reaching the specified age are considered candidates for deduplication",
		"SurvivorRatio=", "Sets the ratio between eden space size and survivor space size",
		"TargetSurvivorRatio=", "Sets the desired percentage of survivor space used after young garbage collection",
		"TLABSize=", "Sets the initial size",
		"+UseAdaptiveSizePolicy", "Enables the use of adaptive generation sizing",
		"+UseCMSInitiatingOccupancyOnly", "Enables the use of the occupancy value as the only criterion for initiating the CMS collector",
		"+UseConcMarkSweepGC", "Enables the use of the CMS garbage collector for the old generation",
		"+UseG1GC", "Enables the use of the garbage-first garbage collector",
		"+UseGCOverheadLimit", "Enables the use of a policy that limits the proportion of time spent by the JVMon GC before an OutOfMemoryError exception is thrown",
		"+UseNUMA", "Enables performance optimization of an application on a machine with nonuniform memory architecture",
		"+UseParallelGC", "Enables the use of the parallel scavenge garbage collector",
		"+UseParallelOldGC", "Enables the use of the parallel garbage collector for full GCs",
		"+UseParNewGC", "Enables the use of parallel threads for collection in the young generation",
		"+UseSerialGC", "Enables the use of the serial garbage collector",
		"+UseSHM", "On Linux, enables the JVM to use shared memory to setup large pages",
		"+UseStringDeduplication", "Enables string deduplication",
		"+UseTLAB", "Enables the use of thread-local allocation blocks in the young generation space",
	)
}

func ActionAdvancedValues(key string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch key {
		case "ErrorFile":
			return carapace.ActionFiles()
		case "LargePageSizeInBytes":
			return carapace.ActionValues()
		case "MaxDirectMemorySize":
			return carapace.ActionValues()
		case "NativeMemoryTracking":
			return carapace.ActionValuesDescribed(
				"off", "Do not track JVM native memory usage",
				"summary", "Only track memory usage by JVM subsystems",
				"detail", "additionally track memory usage by individual CallSite",
			)
		case "ObjectAlignmentInBytes":
			return carapace.ActionValues()
		case "OnError":
			return carapace.ActionValues()
		case "OnOutOfMemoryError":
			return carapace.ActionValues()
		case "ThreadStackSize":
			return carapace.ActionValues()
		case "AllocateInstancePrefetchLines":
			return carapace.ActionValues()
		case "AllocatePrefetchDistance":
			return carapace.ActionValues()
		case "AllocatePrefetchInstr":
			return carapace.ActionValues()
		case "AllocatePrefetchLines":
			return carapace.ActionValues()
		case "AllocatePrefetchStepSize":
			return carapace.ActionValues()
		case "AllocatePrefetchStyle":
			return carapace.ActionValuesDescribed(
				"0", "Do not generate prefetch instructions",
				"1", "Execute prefetch instructions after each allocation",
				"2", "Use the thread-local allocation block (TLAB) watermark pointer to determine when prefetch instructions are executed",
				"3", "Use BIS instruction on SPARC for allocation prefetch",
			)
		case "CICompilerCount":
			return carapace.ActionValues()
		case "CodeCacheMinimumFreeSpace":
			return carapace.ActionValues()
		case "CompileCommand":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"break", "Set a breakpoint when debugging the JVM to stop at the beginning of compilation of the specified method",
						"compileonly", "Exclude all methods from compilation except for the specified method",
						"dontinline", "Prevent inlining of the specified method",
						"exclude", "Exclude the specified method from compilation",
						"help", "Print a help message for the -XX:CompileCommand option",
						"inline", "Attempt to inline the specified method",
						"log", "Exclude compilation logging for all methods except for the specified method",
						"option", "This command can be used to pass a JIT compilation option to the specified method",
						"print", "Print generated assembler code after compilation of the specified method",
						"quiet", "Do not print the compile commands",
					)
				default:
					return carapace.ActionValues()
				}
			})
		case "CompileCommandFile":
			return carapace.ActionFiles()
		case "CompileOnly":
			return carapace.ActionValues()
		case "CompileThreshold":
			return carapace.ActionValues()
		case "InitialCodeCacheSize":
			return carapace.ActionValues()
		case "InlineSmallCode":
			return carapace.ActionValues()
		case "MaxInlineSize":
			return carapace.ActionValues()
		case "MaxNodeLimit":
			return carapace.ActionValues()
		case "MaxTrivialSize":
			return carapace.ActionValues()
		case "ReservedCodeCacheSize":
			return carapace.ActionValues()
		case "RTMAbortRatio":
			return carapace.ActionValues()
		case "RTMRetryCount":
			return carapace.ActionValues()
		case "HeapDumpPath":
			return carapace.ActionFiles()
		case "LogFile":
			return carapace.ActionFiles()
		case "CMSExpAvgFactor":
			return carapace.ActionValues()
		case "CMSInitiatingOccupancyFraction":
			return carapace.ActionValues()
		case "CMSTriggerRatio":
			return carapace.ActionValues()
		case "ConcGCThreads":
			return carapace.ActionValues()
		case "G1HeapRegionSize":
			return carapace.ActionValues()
		case "G1ReservePercent":
			return carapace.ActionValues()
		case "InitialHeapSize":
			return carapace.ActionValues()
		case "InitialSurvivorRatio":
			return carapace.ActionValues()
		case "InitiatingHeapOccupancyPercent":
			return carapace.ActionValues()
		case "MaxGCPauseMillis":
			return carapace.ActionValues()
		case "MaxHeapSize":
			return carapace.ActionValues()
		case "MaxHeapFreeRatio":
			return carapace.ActionValues()
		case "MaxMetaspaceSize":
			return carapace.ActionValues()
		case "MaxNewSize":
			return carapace.ActionValues()
		case "MaxTenuringThreshold":
			return carapace.ActionValues()
		case "MetaspaceSize":
			return carapace.ActionValues()
		case "MinHeapFreeRatio":
			return carapace.ActionValues()
		case "NewRatio":
			return carapace.ActionValues()
		case "NewSize":
			return carapace.ActionValues()
		case "ParallelGCThreads":
			return carapace.ActionValues()
		case "SoftRefLRUPolicyMSPerMB":
			return carapace.ActionValues()
		case "StringDeduplicationAgeThreshold":
			return carapace.ActionValues()
		case "SurvivorRatio":
			return carapace.ActionValues()
		case "TargetSurvivorRatio":
			return carapace.ActionValues()
		case "TLABSize":
			return carapace.ActionValues()
		default:
			return carapace.ActionValues()
		}
	})
}
