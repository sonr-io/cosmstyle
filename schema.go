package nebula

import "go.mongodb.org/mongo-driver/bson"

var schema = bson.M{
	"$schema":            "https://raw.githubusercontent.com/JetBrains/web-types/master/schema/web-types.json",
	"name":               "@onsonr/nebula",
	"version":            "0.0.14",
	"description-markup": "markdown",
	"contributions": bson.M{
		"html": bson.M{
			"elements": bson.A{
				bson.M{
					"name":        "sl-alert",
					"description": "Alerts are used to display important messages inline or as toast notifications.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the alert opens.\n- **sl-after-show** - Emitted after the alert opens and all animations are complete.\n- **sl-hide** - Emitted when the alert closes.\n- **sl-after-hide** - Emitted after the alert closes and all animations are complete.\n\n### **Methods:**\n - **show()** - Shows the alert.\n- **hide()** - Hides the alert\n- **toast()** - Displays the alert as a toast notification. This will move the alert out of its position in the DOM and, when\ndismissed, it will be removed from the DOM completely. By storing a reference to the alert, you can reuse it by\ncalling this method again. The returned promise will resolve after the alert is hidden.\n\n### **Slots:**\n - _default_ - The alert's main content.\n- **icon** - An icon to show in the alert. Works best with `<sl-icon>`.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **icon** - The container that wraps the optional icon.\n- **message** - The container that wraps the alert's main content.\n- **close-button** - The close button, an `<sl-icon-button>`.\n- **close-button__base** - The close button's exported `base` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the alert is open. You can toggle this attribute to show and hide the alert, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the alert's open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "closable",
							"description": "Enables a close button that allows the user to dismiss the alert.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "variant",
							"description": "The alert's theme variant.",
							"value": bson.M{
								"type":    "'primary' | 'success' | 'neutral' | 'warning' | 'danger'",
								"default": "'primary'",
							},
						},
						bson.M{
							"name":        "duration",
							"description": "The length of time, in milliseconds, the alert will show before closing itself. If the user interacts with\nthe alert before it closes (e.g. moves the mouse over it), the timer will restart. Defaults to `Infinity`, meaning\nthe alert will not close on its own.",
							"value": bson.M{
								"type":    "string",
								"default": "Infinity",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The alert's main content.",
						},
						bson.M{
							"name":        "icon",
							"description": "An icon to show in the alert. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the alert opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the alert opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the alert closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the alert closes and all animations are complete.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "base",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the alert is open. You can toggle this attribute to show and hide the alert, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the alert's open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "closable",
								"description": "Enables a close button that allows the user to dismiss the alert.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "variant",
								"description": "The alert's theme variant.",
								"type":        "'primary' | 'success' | 'neutral' | 'warning' | 'danger'",
							},
							bson.M{
								"name":        "duration",
								"description": "The length of time, in milliseconds, the alert will show before closing itself. If the user interacts with\nthe alert before it closes (e.g. moves the mouse over it), the timer will restart. Defaults to `Infinity`, meaning\nthe alert will not close on its own.",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the alert opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the alert opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the alert closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the alert closes and all animations are complete.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-animated-image",
					"description": "A component for displaying animated GIFs and WEBPs that play and pause on interaction.\n---\n\n\n### **Events:**\n - **sl-load** - Emitted when the image loads successfully.\n- **sl-error** - Emitted when the image fails to load.\n\n### **Slots:**\n - **play-icon** - Optional play icon to use instead of the default. Works best with `<sl-icon>`.\n- **pause-icon** - Optional pause icon to use instead of the default. Works best with `<sl-icon>`.\n\n### **CSS Properties:**\n - **--control-box-size** - The size of the icon box. _(default: undefined)_\n- **--icon-size** - The size of the play/pause icons. _(default: undefined)_\n\n### **CSS Parts:**\n - **control-box** - The container that surrounds the pause/play icons and provides their background.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "src",
							"description": "The path to the image to load.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "alt",
							"description": "A description of the image used by assistive devices.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "play",
							"description": "Plays the animation. When this attribute is remove, the animation will pause.",
							"value": bson.M{
								"type": "boolean",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "play-icon",
							"description": "Optional play icon to use instead of the default. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "pause-icon",
							"description": "Optional pause icon to use instead of the default. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-load",
							"description": "Emitted when the image loads successfully.",
						},
						bson.M{
							"name":        "sl-error",
							"description": "Emitted when the image fails to load.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "animatedImage",
								"type": "HTMLImageElement",
							},
							bson.M{
								"name": "frozenFrame",
								"type": "string",
							},
							bson.M{
								"name": "isLoaded",
								"type": "boolean",
							},
							bson.M{
								"name":        "src",
								"description": "The path to the image to load.",
								"type":        "string",
							},
							bson.M{
								"name":        "alt",
								"description": "A description of the image used by assistive devices.",
								"type":        "string",
							},
							bson.M{
								"name":        "play",
								"description": "Plays the animation. When this attribute is remove, the animation will pause.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-load",
								"description": "Emitted when the image loads successfully.",
							},
							bson.M{
								"name":        "sl-error",
								"description": "Emitted when the image fails to load.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-animation",
					"description": "Animate elements declaratively with nearly 100 baked-in presets, or roll your own with custom keyframes. Powered by the bson.A{Web Animations API}(https://developer.mozilla.org/en-US/docs/Web/API/Web_Animations_API).\n---\n\n\n### **Events:**\n - **sl-cancel** - Emitted when the animation is canceled.\n- **sl-finish** - Emitted when the animation finishes.\n- **sl-start** - Emitted when the animation starts or restarts.\n\n### **Methods:**\n - **cancel()** - Clears all keyframe effects caused by this animation and aborts its playback.\n- **finish()** - Sets the playback time to the end of the animation corresponding to the current playback direction.\n\n### **Slots:**\n - _default_ - The element to animate. Avoid slotting in more than one element, as subsequent ones will be ignored. To animate multiple elements, either wrap them in a single container or use multiple `<sl-animation>` elements.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "name",
							"description": "The name of the built-in animation to use. For custom animations, use the `keyframes` prop.",
							"value": bson.M{
								"type":    "string",
								"default": "'none'",
							},
						},
						bson.M{
							"name":        "play",
							"description": "Plays the animation. When omitted, the animation will be paused. This attribute will be automatically removed when\nthe animation finishes or gets canceled.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "delay",
							"description": "The number of milliseconds to delay the start of the animation.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "direction",
							"description": "Determines the direction of playback as well as the behavior when reaching the end of an iteration.\nbson.A{Learn more}(https://developer.mozilla.org/en-US/docs/Web/CSS/animation-direction)",
							"value": bson.M{
								"type":    "PlaybackDirection",
								"default": "'normal'",
							},
						},
						bson.M{
							"name":        "duration",
							"description": "The number of milliseconds each iteration of the animation takes to complete.",
							"value": bson.M{
								"type":    "number",
								"default": "1000",
							},
						},
						bson.M{
							"name":        "easing",
							"description": "The easing function to use for the animation. This can be a Shoelace easing function or a custom easing function\nsuch as `cubic-bezier(0, 1, .76, 1.14)`.",
							"value": bson.M{
								"type":    "string",
								"default": "'linear'",
							},
						},
						bson.M{
							"name":        "end-delay",
							"description": "The number of milliseconds to delay after the active period of an animation sequence.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "fill",
							"description": "Sets how the animation applies styles to its target before and after its execution.",
							"value": bson.M{
								"type":    "FillMode",
								"default": "'auto'",
							},
						},
						bson.M{
							"name":        "iterations",
							"description": "The number of iterations to run before the animation completes. Defaults to `Infinity`, which loops.",
							"value": bson.M{
								"type":    "string",
								"default": "Infinity",
							},
						},
						bson.M{
							"name":        "iteration-start",
							"description": "The offset at which to start the animation, usually between 0 (start) and 1 (end).",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "playback-rate",
							"description": "Sets the animation's playback rate. The default is `1`, which plays the animation at a normal speed. Setting this\nto `2`, for example, will double the animation's speed. A negative value can be used to reverse the animation. This\nvalue can be changed without causing the animation to restart.",
							"value": bson.M{
								"type":    "number",
								"default": "1",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The element to animate. Avoid slotting in more than one element, as subsequent ones will be ignored. To animate multiple elements, either wrap them in a single container or use multiple `<sl-animation>` elements.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-cancel",
							"description": "Emitted when the animation is canceled.",
						},
						bson.M{
							"name":        "sl-finish",
							"description": "Emitted when the animation finishes.",
						},
						bson.M{
							"name":        "sl-start",
							"description": "Emitted when the animation starts or restarts.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "Promise<HTMLSlotElement>",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the built-in animation to use. For custom animations, use the `keyframes` prop.",
								"type":        "string",
							},
							bson.M{
								"name":        "play",
								"description": "Plays the animation. When omitted, the animation will be paused. This attribute will be automatically removed when\nthe animation finishes or gets canceled.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "delay",
								"description": "The number of milliseconds to delay the start of the animation.",
								"type":        "number",
							},
							bson.M{
								"name":        "direction",
								"description": "Determines the direction of playback as well as the behavior when reaching the end of an iteration.\nbson.A{Learn more}(https://developer.mozilla.org/en-US/docs/Web/CSS/animation-direction)",
								"type":        "PlaybackDirection",
							},
							bson.M{
								"name":        "duration",
								"description": "The number of milliseconds each iteration of the animation takes to complete.",
								"type":        "number",
							},
							bson.M{
								"name":        "easing",
								"description": "The easing function to use for the animation. This can be a Shoelace easing function or a custom easing function\nsuch as `cubic-bezier(0, 1, .76, 1.14)`.",
								"type":        "string",
							},
							bson.M{
								"name":        "endDelay",
								"description": "The number of milliseconds to delay after the active period of an animation sequence.",
								"type":        "number",
							},
							bson.M{
								"name":        "fill",
								"description": "Sets how the animation applies styles to its target before and after its execution.",
								"type":        "FillMode",
							},
							bson.M{
								"name":        "iterations",
								"description": "The number of iterations to run before the animation completes. Defaults to `Infinity`, which loops.",
							},
							bson.M{
								"name":        "iterationStart",
								"description": "The offset at which to start the animation, usually between 0 (start) and 1 (end).",
								"type":        "number",
							},
							bson.M{
								"name":        "keyframes",
								"description": "The keyframes to use for the animation. If this is set, `name` will be ignored.",
								"type":        "Keyframebson.A{} | undefined",
							},
							bson.M{
								"name":        "playbackRate",
								"description": "Sets the animation's playback rate. The default is `1`, which plays the animation at a normal speed. Setting this\nto `2`, for example, will double the animation's speed. A negative value can be used to reverse the animation. This\nvalue can be changed without causing the animation to restart.",
								"type":        "number",
							},
							bson.M{
								"name":        "currentTime",
								"description": "Gets and sets the current animation time.",
								"type":        "CSSNumberish",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-cancel",
								"description": "Emitted when the animation is canceled.",
							},
							bson.M{
								"name":        "sl-finish",
								"description": "Emitted when the animation finishes.",
							},
							bson.M{
								"name":        "sl-start",
								"description": "Emitted when the animation starts or restarts.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-avatar",
					"description": "Avatars are used to represent a person or object.\n---\n\n\n### **Slots:**\n - **icon** - The default icon to use when no image or initials are present. Works best with `<sl-icon>`.\n\n### **CSS Properties:**\n - **--size** - The size of the avatar. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **icon** - The container that wraps the avatar's icon.\n- **initials** - The container that wraps the avatar's initials.\n- **image** - The avatar image. Only shown when the `image` attribute is set.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "image",
							"description": "The image source to use for the avatar.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "label",
							"description": "A label to use to describe the avatar to assistive devices.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "initials",
							"description": "Initials to use as a fallback when no image is available (1-2 characters max recommended).",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "loading",
							"description": "Indicates how the browser should load the image.",
							"value": bson.M{
								"type":    "'eager' | 'lazy'",
								"default": "'eager'",
							},
						},
						bson.M{
							"name":        "shape",
							"description": "The shape of the avatar.",
							"value": bson.M{
								"type":    "'circle' | 'square' | 'rounded'",
								"default": "'circle'",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "icon",
							"description": "The default icon to use when no image or initials are present. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "image",
								"description": "The image source to use for the avatar.",
								"type":        "string",
							},
							bson.M{
								"name":        "label",
								"description": "A label to use to describe the avatar to assistive devices.",
								"type":        "string",
							},
							bson.M{
								"name":        "initials",
								"description": "Initials to use as a fallback when no image is available (1-2 characters max recommended).",
								"type":        "string",
							},
							bson.M{
								"name":        "loading",
								"description": "Indicates how the browser should load the image.",
								"type":        "'eager' | 'lazy'",
							},
							bson.M{
								"name":        "shape",
								"description": "The shape of the avatar.",
								"type":        "'circle' | 'square' | 'rounded'",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-badge",
					"description": "Badges are used to draw attention and display statuses or counts.\n---\n\n\n### **Slots:**\n - _default_ - The badge's content.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "variant",
							"description": "The badge's theme variant.",
							"value": bson.M{
								"type":    "'primary' | 'success' | 'neutral' | 'warning' | 'danger'",
								"default": "'primary'",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style badge with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pulse",
							"description": "Makes the badge pulsate to draw attention.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The badge's content.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "variant",
								"description": "The badge's theme variant.",
								"type":        "'primary' | 'success' | 'neutral' | 'warning' | 'danger'",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style badge with rounded edges.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pulse",
								"description": "Makes the badge pulsate to draw attention.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-breadcrumb-item",
					"description": "Breadcrumb Items are used inside bson.A{breadcrumbs}(/components/breadcrumb) to represent different links.\n---\n\n\n### **Slots:**\n - _default_ - The breadcrumb item's label.\n- **prefix** - An optional prefix, usually an icon or icon button.\n- **suffix** - An optional suffix, usually an icon or icon button.\n- **separator** - The separator to use for the breadcrumb item. This will only change the separator for this item. If you want to change it for all items in the group, set the separator on `<sl-breadcrumb>` instead.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **label** - The breadcrumb item's label.\n- **prefix** - The container that wraps the prefix.\n- **suffix** - The container that wraps the suffix.\n- **separator** - The container that wraps the separator.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "href",
							"description": "Optional URL to direct the user to when the breadcrumb item is activated. When set, a link will be rendered\ninternally. When unset, a button will be rendered instead.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "target",
							"description": "Tells the browser where to open the link. Only used when `href` is set.",
							"value": bson.M{
								"type": "'_blank' | '_parent' | '_self' | '_top' | undefined",
							},
						},
						bson.M{
							"name":        "rel",
							"description": "The `rel` attribute to use on the link. Only used when `href` is set.",
							"value": bson.M{
								"type":    "string",
								"default": "'noreferrer noopener'",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The breadcrumb item's label.",
						},
						bson.M{
							"name":        "prefix",
							"description": "An optional prefix, usually an icon or icon button.",
						},
						bson.M{
							"name":        "suffix",
							"description": "An optional suffix, usually an icon or icon button.",
						},
						bson.M{
							"name":        "separator",
							"description": "The separator to use for the breadcrumb item. This will only change the separator for this item. If you want to change it for all items in the group, set the separator on `<sl-breadcrumb>` instead.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "href",
								"description": "Optional URL to direct the user to when the breadcrumb item is activated. When set, a link will be rendered\ninternally. When unset, a button will be rendered instead.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "target",
								"description": "Tells the browser where to open the link. Only used when `href` is set.",
								"type":        "'_blank' | '_parent' | '_self' | '_top' | undefined",
							},
							bson.M{
								"name":        "rel",
								"description": "The `rel` attribute to use on the link. Only used when `href` is set.",
								"type":        "string",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-breadcrumb",
					"description": "Breadcrumbs provide a group of links so users can easily navigate a website's hierarchy.\n---\n\n\n### **Slots:**\n - _default_ - One or more breadcrumb items to display.\n- **separator** - The separator to use between breadcrumb items. Works best with `<sl-icon>`.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "label",
							"description": "The label to use for the breadcrumb control. This will not be shown on the screen, but it will be announced by\nscreen readers and other assistive devices to provide more context for users.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "One or more breadcrumb items to display.",
						},
						bson.M{
							"name":        "separator",
							"description": "The separator to use between breadcrumb items. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "separatorSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name":        "label",
								"description": "The label to use for the breadcrumb control. This will not be shown on the screen, but it will be announced by\nscreen readers and other assistive devices to provide more context for users.",
								"type":        "string",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-button",
					"description": "Buttons represent actions that are available to the user.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the button loses focus.\n- **sl-focus** - Emitted when the button gains focus.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **click()** - Simulates a click on the button.\n- **focus(options: _FocusOptions_)** - Sets focus on the button.\n- **blur()** - Removes focus from the button.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - _default_ - The button's label.\n- **prefix** - A presentational prefix icon or similar element.\n- **suffix** - A presentational suffix icon or similar element.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **prefix** - The container that wraps the prefix.\n- **label** - The button's label.\n- **suffix** - The container that wraps the suffix.\n- **caret** - The button's caret icon, an `<sl-icon>` element.\n- **spinner** - The spinner that shows when the button is in the loading state.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "variant",
							"description": "The button's theme variant.",
							"value": bson.M{
								"type":    "'default' | 'primary' | 'success' | 'neutral' | 'warning' | 'danger' | 'text'",
								"default": "'default'",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The button's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "caret",
							"description": "Draws the button with a caret. Used to indicate that the button triggers a dropdown menu or similar behavior.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "loading",
							"description": "Draws the button in a loading state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "outline",
							"description": "Draws an outlined button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style button with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "circle",
							"description": "Draws a circular icon button. When this attribute is present, the button expects a single `<sl-icon>` in the\ndefault slot.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "type",
							"description": "The type of button. Note that the default value is `button` instead of `submit`, which is opposite of how native\n`<button>` elements behave. When the type is `submit`, the button will submit the surrounding form.",
							"value": bson.M{
								"type":    "'button' | 'submit' | 'reset'",
								"default": "'button'",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the button, submitted as a name/value pair with form data, but only when this button is the submitter.\nThis attribute is ignored when `href` is present.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The value of the button, submitted as a pair with the button's name as part of the form data, but only when this\nbutton is the submitter. This attribute is ignored when `href` is present.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "href",
							"description": "When set, the underlying button will be rendered as an `<a>` with this `href` instead of a `<button>`.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "target",
							"description": "Tells the browser where to open the link. Only used when `href` is present.",
							"value": bson.M{
								"type": "'_blank' | '_parent' | '_self' | '_top'",
							},
						},
						bson.M{
							"name":        "rel",
							"description": "When using `href`, this attribute will map to the underlying link's `rel` attribute. Unlike regular links, the\ndefault is `noreferrer noopener` to prevent security exploits. However, if you're using `target` to point to a\nspecific tab/window, this will prevent that from working correctly. You can remove or change the default value by\nsetting the attribute to an empty string or a value of your choice, respectively.",
							"value": bson.M{
								"type":    "string",
								"default": "'noreferrer noopener'",
							},
						},
						bson.M{
							"name":        "download",
							"description": "Tells the browser to download the linked file as this filename. Only used when `href` is present.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "form",
							"description": "The \"form owner\" to associate the button with. If omitted, the closest containing form will be used instead. The\nvalue of this attribute must be an id of a form in the same document or shadow root as the button.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "formaction",
							"description": "Used to override the form owner's `action` attribute.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "formenctype",
							"description": "Used to override the form owner's `enctype` attribute.",
							"value": bson.M{
								"type": "'application/x-www-form-urlencoded' | 'multipart/form-data' | 'text/plain'",
							},
						},
						bson.M{
							"name":        "formmethod",
							"description": "Used to override the form owner's `method` attribute.",
							"value": bson.M{
								"type": "'post' | 'get'",
							},
						},
						bson.M{
							"name":        "formnovalidate",
							"description": "Used to override the form owner's `novalidate` attribute.",
							"value": bson.M{
								"type": "boolean",
							},
						},
						bson.M{
							"name":        "formtarget",
							"description": "Used to override the form owner's `target` attribute.",
							"value": bson.M{
								"type": "'_self' | '_blank' | '_parent' | '_top' | string",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The button's label.",
						},
						bson.M{
							"name":        "prefix",
							"description": "A presentational prefix icon or similar element.",
						},
						bson.M{
							"name":        "suffix",
							"description": "A presentational suffix icon or similar element.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the button loses focus.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the button gains focus.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "button",
								"type": "HTMLButtonElement | HTMLLinkElement",
							},
							bson.M{
								"name": "invalid",
								"type": "boolean",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "variant",
								"description": "The button's theme variant.",
								"type":        "'default' | 'primary' | 'success' | 'neutral' | 'warning' | 'danger' | 'text'",
							},
							bson.M{
								"name":        "size",
								"description": "The button's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "caret",
								"description": "Draws the button with a caret. Used to indicate that the button triggers a dropdown menu or similar behavior.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the button.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "loading",
								"description": "Draws the button in a loading state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "outline",
								"description": "Draws an outlined button.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style button with rounded edges.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "circle",
								"description": "Draws a circular icon button. When this attribute is present, the button expects a single `<sl-icon>` in the\ndefault slot.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "type",
								"description": "The type of button. Note that the default value is `button` instead of `submit`, which is opposite of how native\n`<button>` elements behave. When the type is `submit`, the button will submit the surrounding form.",
								"type":        "'button' | 'submit' | 'reset'",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the button, submitted as a name/value pair with form data, but only when this button is the submitter.\nThis attribute is ignored when `href` is present.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The value of the button, submitted as a pair with the button's name as part of the form data, but only when this\nbutton is the submitter. This attribute is ignored when `href` is present.",
								"type":        "string",
							},
							bson.M{
								"name":        "href",
								"description": "When set, the underlying button will be rendered as an `<a>` with this `href` instead of a `<button>`.",
								"type":        "string",
							},
							bson.M{
								"name":        "target",
								"description": "Tells the browser where to open the link. Only used when `href` is present.",
								"type":        "'_blank' | '_parent' | '_self' | '_top'",
							},
							bson.M{
								"name":        "rel",
								"description": "When using `href`, this attribute will map to the underlying link's `rel` attribute. Unlike regular links, the\ndefault is `noreferrer noopener` to prevent security exploits. However, if you're using `target` to point to a\nspecific tab/window, this will prevent that from working correctly. You can remove or change the default value by\nsetting the attribute to an empty string or a value of your choice, respectively.",
								"type":        "string",
							},
							bson.M{
								"name":        "download",
								"description": "Tells the browser to download the linked file as this filename. Only used when `href` is present.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "form",
								"description": "The \"form owner\" to associate the button with. If omitted, the closest containing form will be used instead. The\nvalue of this attribute must be an id of a form in the same document or shadow root as the button.",
								"type":        "string",
							},
							bson.M{
								"name":        "formAction",
								"description": "Used to override the form owner's `action` attribute.",
								"type":        "string",
							},
							bson.M{
								"name":        "formEnctype",
								"description": "Used to override the form owner's `enctype` attribute.",
								"type":        "'application/x-www-form-urlencoded' | 'multipart/form-data' | 'text/plain'",
							},
							bson.M{
								"name":        "formMethod",
								"description": "Used to override the form owner's `method` attribute.",
								"type":        "'post' | 'get'",
							},
							bson.M{
								"name":        "formNoValidate",
								"description": "Used to override the form owner's `novalidate` attribute.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "formTarget",
								"description": "Used to override the form owner's `target` attribute.",
								"type":        "'_self' | '_blank' | '_parent' | '_top' | string",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the button loses focus.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the button gains focus.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-button-group",
					"description": "Button groups can be used to group related buttons into sections.\n---\n\n\n### **Slots:**\n - _default_ - One or more `<sl-button>` elements to display in the button group.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "label",
							"description": "A label to use for the button group. This won't be displayed on the screen, but it will be announced by assistive\ndevices when interacting with the control and is strongly recommended.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "One or more `<sl-button>` elements to display in the button group.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "disableRole",
								"type": "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "A label to use for the button group. This won't be displayed on the screen, but it will be announced by assistive\ndevices when interacting with the control and is strongly recommended.",
								"type":        "string",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-card",
					"description": "Cards can be used to group related subjects in a container.\n---\n\n\n### **Slots:**\n - _default_ - The card's main content.\n- **header** - An optional header for the card.\n- **footer** - An optional footer for the card.\n- **image** - An optional image to render at the start of the card.\n\n### **CSS Properties:**\n - **--border-color** - The card's border color, including borders that occur inside the card. _(default: undefined)_\n- **--border-radius** - The border radius for the card's edges. _(default: undefined)_\n- **--border-width** - The width of the card's borders. _(default: undefined)_\n- **--padding** - The padding to use for the card's sections. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **image** - The container that wraps the card's image.\n- **header** - The container that wraps the card's header.\n- **body** - The container that wraps the card's main content.\n- **footer** - The container that wraps the card's footer.",
					"doc-url":     "",
					"attributes":  bson.A{},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The card's main content.",
						},
						bson.M{
							"name":        "header",
							"description": "An optional header for the card.",
						},
						bson.M{
							"name":        "footer",
							"description": "An optional footer for the card.",
						},
						bson.M{
							"name":        "image",
							"description": "An optional image to render at the start of the card.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{},
						"events":     bson.A{},
					},
				},
				bson.M{
					"name":        "sl-carousel",
					"description": "Carousels display an arbitrary number of content slides along a horizontal or vertical axis.\n---\n\n\n### **Events:**\n - **sl-slide-change** - Emitted when the active slide changes.\n\n### **Methods:**\n - **previous(behavior: _ScrollBehavior_)** - Move the carousel backward by `slides-per-move` slides.\n- **next(behavior: _ScrollBehavior_)** - Move the carousel forward by `slides-per-move` slides.\n- **goToSlide(index: _number_, behavior: _ScrollBehavior_)** - Scrolls the carousel to the slide specified by `index`.\n\n### **Slots:**\n - _default_ - The carousel's main content, one or more `<sl-carousel-item>` elements.\n- **next-icon** - Optional next icon to use instead of the default. Works best with `<sl-icon>`.\n- **previous-icon** - Optional previous icon to use instead of the default. Works best with `<sl-icon>`.\n\n### **CSS Properties:**\n - **--slide-gap** - The space between each slide. _(default: undefined)_\n- **--aspect-ratio** - The aspect ratio of each slide. _(default: 16/9)_\n- **--scroll-hint** - The amount of padding to apply to the scroll area, allowing adjacent slides to become partially visible as a scroll hint. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The carousel's internal wrapper.\n- **scroll-container** - The scroll container that wraps the slides.\n- **pagination** - The pagination indicators wrapper.\n- **pagination-item** - The pagination indicator.\n- **pagination-item--active** - Applied when the item is active.\n- **navigation** - The navigation wrapper.\n- **navigation-button** - The navigation button.\n- **navigation-button--previous** - Applied to the previous button.\n- **navigation-button--next** - Applied to the next button.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "loop",
							"description": "When set, allows the user to navigate the carousel in the same direction indefinitely.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "navigation",
							"description": "When set, show the carousel's navigation.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pagination",
							"description": "When set, show the carousel's pagination indicators.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "autoplay",
							"description": "When set, the slides will scroll automatically when the user is not interacting with them.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "autoplay-interval",
							"description": "Specifies the amount of time, in milliseconds, between each automatic scroll.",
							"value": bson.M{
								"type":    "number",
								"default": "3000",
							},
						},
						bson.M{
							"name":        "slides-per-page",
							"description": "Specifies how many slides should be shown at a given time.",
							"value": bson.M{
								"type":    "number",
								"default": "1",
							},
						},
						bson.M{
							"name":        "slides-per-move",
							"description": "Specifies the number of slides the carousel will advance when scrolling, useful when specifying a `slides-per-page`\ngreater than one. It can't be higher than `slides-per-page`.",
							"value": bson.M{
								"type":    "number",
								"default": "1",
							},
						},
						bson.M{
							"name":        "orientation",
							"description": "Specifies the orientation in which the carousel will lay out.",
							"value": bson.M{
								"type":    "'horizontal' | 'vertical'",
								"default": "'horizontal'",
							},
						},
						bson.M{
							"name":        "mouse-dragging",
							"description": "When set, it is possible to scroll through the slides by dragging them with the mouse.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The carousel's main content, one or more `<sl-carousel-item>` elements.",
						},
						bson.M{
							"name":        "next-icon",
							"description": "Optional next icon to use instead of the default. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "previous-icon",
							"description": "Optional previous icon to use instead of the default. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-slide-change",
							"type":        "bson.M{ index: number, slide: SlCarouselItem }",
							"description": "Emitted when the active slide changes.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "loop",
								"description": "When set, allows the user to navigate the carousel in the same direction indefinitely.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "navigation",
								"description": "When set, show the carousel's navigation.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pagination",
								"description": "When set, show the carousel's pagination indicators.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "autoplay",
								"description": "When set, the slides will scroll automatically when the user is not interacting with them.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "autoplayInterval",
								"description": "Specifies the amount of time, in milliseconds, between each automatic scroll.",
								"type":        "number",
							},
							bson.M{
								"name":        "slidesPerPage",
								"description": "Specifies how many slides should be shown at a given time.",
								"type":        "number",
							},
							bson.M{
								"name":        "slidesPerMove",
								"description": "Specifies the number of slides the carousel will advance when scrolling, useful when specifying a `slides-per-page`\ngreater than one. It can't be higher than `slides-per-page`.",
								"type":        "number",
							},
							bson.M{
								"name":        "orientation",
								"description": "Specifies the orientation in which the carousel will lay out.",
								"type":        "'horizontal' | 'vertical'",
							},
							bson.M{
								"name":        "mouseDragging",
								"description": "When set, it is possible to scroll through the slides by dragging them with the mouse.",
								"type":        "boolean",
							},
							bson.M{
								"name": "scrollContainer",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "paginationContainer",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "activeSlide",
								"type": "number",
							},
							bson.M{
								"name": "scrolling",
								"type": "boolean",
							},
							bson.M{
								"name": "dragging",
								"type": "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-slide-change",
								"type":        "bson.M{ index: number, slide: SlCarouselItem }",
								"description": "Emitted when the active slide changes.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-carousel-item",
					"description": "A carousel item represent a slide within a bson.A{carousel}(/components/carousel).\n---\n\n\n### **Slots:**\n - _default_ - The carousel item's content..\n\n### **CSS Properties:**\n - **--aspect-ratio** - The slide's aspect ratio. Inherited from the carousel by default. _(default: undefined)_",
					"doc-url":     "",
					"attributes":  bson.A{},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The carousel item's content..",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{},
						"events":     bson.A{},
					},
				},
				bson.M{
					"name":        "sl-checkbox",
					"description": "Checkboxes allow the user to toggle an option on or off.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the checkbox loses focus.\n- **sl-change** - Emitted when the checked state changes.\n- **sl-focus** - Emitted when the checkbox gains focus.\n- **sl-input** - Emitted when the checkbox receives input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **click()** - Simulates a click on the checkbox.\n- **focus(options: _FocusOptions_)** - Sets focus on the checkbox.\n- **blur()** - Removes focus from the checkbox.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. The value provided will be shown to the user when the form is submitted. To clear\nthe custom validation message, call this method with an empty string.\n\n### **Slots:**\n - _default_ - The checkbox's label.\n- **help-text** - Text that describes how to use the checkbox. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **control** - The square container that wraps the checkbox's checked state.\n- **control--checked** - Matches the control part when the checkbox is checked.\n- **control--indeterminate** - Matches the control part when the checkbox is indeterminate.\n- **checked-icon** - The checked icon, an `<sl-icon>` element.\n- **indeterminate-icon** - The indeterminate icon, an `<sl-icon>` element.\n- **label** - The container that wraps the checkbox's label.\n- **form-control-help-text** - The help text's wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the checkbox, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the checkbox, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The checkbox's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the checkbox.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "checked",
							"description": "Draws the checkbox in a checked state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "indeterminate",
							"description": "Draws the checkbox in an indeterminate state. This is usually applied to checkboxes that represents a \"select\nall/none\" behavior when associated checkboxes have a mix of checked and unchecked states.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Makes the checkbox a required field.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The checkbox's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The checkbox's label.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the checkbox. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the checkbox loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the checked state changes.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the checkbox gains focus.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the checkbox receives input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the checkbox, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the checkbox, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The checkbox's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the checkbox.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "checked",
								"description": "Draws the checkbox in a checked state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "indeterminate",
								"description": "Draws the checkbox in an indeterminate state. This is usually applied to checkboxes that represents a \"select\nall/none\" behavior when associated checkboxes have a mix of checked and unchecked states.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "defaultChecked",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Makes the checkbox a required field.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "helpText",
								"description": "The checkbox's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the checkbox loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the checked state changes.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the checkbox gains focus.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the checkbox receives input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-color-picker",
					"description": "Color pickers allow the user to select a color.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the color picker loses focus.\n- **sl-change** - Emitted when the color picker's value changes.\n- **sl-focus** - Emitted when the color picker receives focus.\n- **sl-input** - Emitted when the color picker receives input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the color picker.\n- **blur()** - Removes focus from the color picker.\n- **getFormattedValue(format: _'hex' | 'hexa' | 'rgb' | 'rgba' | 'hsl' | 'hsla' | 'hsv' | 'hsva'_)** - Returns the current value as a string in the specified format.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - **label** - The color picker's form label. Alternatively, you can use the `label` attribute.\n\n### **CSS Properties:**\n - **--grid-width** - The width of the color grid. _(default: undefined)_\n- **--grid-height** - The height of the color grid. _(default: undefined)_\n- **--grid-handle-size** - The size of the color grid's handle. _(default: undefined)_\n- **--slider-height** - The height of the hue and alpha sliders. _(default: undefined)_\n- **--slider-handle-size** - The diameter of the slider's handle. _(default: undefined)_\n- **--swatch-size** - The size of each predefined color swatch. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **trigger** - The color picker's dropdown trigger.\n- **swatches** - The container that holds the swatches.\n- **swatch** - Each individual swatch.\n- **grid** - The color grid.\n- **grid-handle** - The color grid's handle.\n- **slider** - Hue and opacity sliders.\n- **slider-handle** - Hue and opacity slider handles.\n- **hue-slider** - The hue slider.\n- **hue-slider-handle** - The hue slider's handle.\n- **opacity-slider** - The opacity slider.\n- **opacity-slider-handle** - The opacity slider's handle.\n- **preview** - The preview color.\n- **input** - The text input.\n- **eye-dropper-button** - The eye dropper button.\n- **eye-dropper-button__base** - The eye dropper button's exported `button` part.\n- **eye-dropper-button__prefix** - The eye dropper button's exported `prefix` part.\n- **eye-dropper-button__label** - The eye dropper button's exported `label` part.\n- **eye-dropper-button__suffix** - The eye dropper button's exported `suffix` part.\n- **eye-dropper-button__caret** - The eye dropper button's exported `caret` part.\n- **format-button** - The format button.\n- **format-button__base** - The format button's exported `button` part.\n- **format-button__prefix** - The format button's exported `prefix` part.\n- **format-button__label** - The format button's exported `label` part.\n- **format-button__suffix** - The format button's exported `suffix` part.\n- **format-button__caret** - The format button's exported `caret` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The current value of the color picker. The value's format will vary based the `format` attribute. To get the value\nin a specific format, use the `getFormattedValue()` method. The value is submitted as a name/value pair with form\ndata.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The color picker's label. This will not be displayed, but it will be announced by assistive devices. If you need to\ndisplay HTML, you can use the `label` slot` instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "format",
							"description": "The format to use. If opacity is enabled, these will translate to HEXA, RGBA, HSLA, and HSVA respectively. The color\npicker will accept user input in any format (including CSS color names) and convert it to the desired format.",
							"value": bson.M{
								"type":    "'hex' | 'rgb' | 'hsl' | 'hsv'",
								"default": "'hex'",
							},
						},
						bson.M{
							"name":        "inline",
							"description": "Renders the color picker inline rather than in a dropdown.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "size",
							"description": "Determines the size of the color picker's trigger. This has no effect on inline color pickers.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "no-format-toggle",
							"description": "Removes the button that lets users toggle between format.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the form control, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the color picker.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "hoist",
							"description": "Enable this option to prevent the panel from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "opacity",
							"description": "Shows the opacity slider. Enabling this will cause the formatted value to be HEXA, RGBA, or HSLA.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "uppercase",
							"description": "By default, values are lowercase. With this attribute, values will be uppercase instead.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "swatches",
							"description": "One or more predefined color swatches to display as presets in the color picker. Can include any format the color\npicker can parse, including HEX(A), RGB(A), HSL(A), HSV(A), and CSS color names. Each color must be separated by a\nsemicolon (`;`). Alternatively, you can pass an array of color values to this property using JavaScript.",
							"value": bson.M{
								"type":    "string | stringbson.A{}",
								"default": "''",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Makes the color picker a required field.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "label",
							"description": "The color picker's form label. Alternatively, you can use the `label` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the color picker loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the color picker's value changes.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the color picker receives focus.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the color picker receives input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "base",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "input",
								"type": "SlInput",
							},
							bson.M{
								"name": "dropdown",
								"type": "SlDropdown",
							},
							bson.M{
								"name": "previewButton",
								"type": "HTMLButtonElement",
							},
							bson.M{
								"name": "trigger",
								"type": "HTMLButtonElement",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the color picker. The value's format will vary based the `format` attribute. To get the value\nin a specific format, use the `getFormattedValue()` method. The value is submitted as a name/value pair with form\ndata.",
								"type":        "string",
							},
							bson.M{
								"name":        "defaultValue",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "string",
							},
							bson.M{
								"name":        "label",
								"description": "The color picker's label. This will not be displayed, but it will be announced by assistive devices. If you need to\ndisplay HTML, you can use the `label` slot` instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "format",
								"description": "The format to use. If opacity is enabled, these will translate to HEXA, RGBA, HSLA, and HSVA respectively. The color\npicker will accept user input in any format (including CSS color names) and convert it to the desired format.",
								"type":        "'hex' | 'rgb' | 'hsl' | 'hsv'",
							},
							bson.M{
								"name":        "inline",
								"description": "Renders the color picker inline rather than in a dropdown.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "size",
								"description": "Determines the size of the color picker's trigger. This has no effect on inline color pickers.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "noFormatToggle",
								"description": "Removes the button that lets users toggle between format.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the form control, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the color picker.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "hoist",
								"description": "Enable this option to prevent the panel from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "opacity",
								"description": "Shows the opacity slider. Enabling this will cause the formatted value to be HEXA, RGBA, or HSLA.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "uppercase",
								"description": "By default, values are lowercase. With this attribute, values will be uppercase instead.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "swatches",
								"description": "One or more predefined color swatches to display as presets in the color picker. Can include any format the color\npicker can parse, including HEX(A), RGB(A), HSL(A), HSV(A), and CSS color names. Each color must be separated by a\nsemicolon (`;`). Alternatively, you can pass an array of color values to this property using JavaScript.",
								"type":        "string | stringbson.A{}",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Makes the color picker a required field.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the color picker loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the color picker's value changes.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the color picker receives focus.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the color picker receives input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-copy-button",
					"description": "Copies text data to the clipboard when the user clicks the trigger.\n---\n\n\n### **Events:**\n - **sl-copy** - Emitted when the data has been copied.\n- **sl-error** - Emitted when the data could not be copied.\n\n### **Slots:**\n - **copy-icon** - The icon to show in the default copy state. Works best with `<sl-icon>`.\n- **success-icon** - The icon to show when the content is copied. Works best with `<sl-icon>`.\n- **error-icon** - The icon to show when a copy error occurs. Works best with `<sl-icon>`.\n\n### **CSS Properties:**\n - **--success-color** - The color to use for success feedback. _(default: undefined)_\n- **--error-color** - The color to use for error feedback. _(default: undefined)_\n\n### **CSS Parts:**\n - **button** - The internal `<button>` element.\n- **copy-icon** - The container that holds the copy icon.\n- **success-icon** - The container that holds the success icon.\n- **error-icon** - The container that holds the error icon.\n- **tooltip__base** - The tooltip's exported `base` part.\n- **tooltip__base__popup** - The tooltip's exported `popup` part.\n- **tooltip__base__arrow** - The tooltip's exported `arrow` part.\n- **tooltip__body** - The tooltip's exported `body` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The text value to copy.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "from",
							"description": "An id that references an element in the same document from which data will be copied. If both this and `value` are\npresent, this value will take precedence. By default, the target element's `textContent` will be copied. To copy an\nattribute, append the attribute name wrapped in square brackets, e.g. `from=\"elbson.A{value}\"`. To copy a property,\nappend a dot and the property name, e.g. `from=\"el.value\"`.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the copy button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "copy-label",
							"description": "A custom label to show in the tooltip.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "success-label",
							"description": "A custom label to show in the tooltip after copying.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "error-label",
							"description": "A custom label to show in the tooltip when a copy error occurs.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "feedback-duration",
							"description": "The length of time to show feedback before restoring the default trigger.",
							"value": bson.M{
								"type":    "number",
								"default": "1000",
							},
						},
						bson.M{
							"name":        "tooltip-placement",
							"description": "The preferred placement of the tooltip.",
							"value": bson.M{
								"type":    "'top' | 'right' | 'bottom' | 'left'",
								"default": "'top'",
							},
						},
						bson.M{
							"name":        "hoist",
							"description": "Enable this option to prevent the tooltip from being clipped when the component is placed inside a container with\n`overflow: auto|hidden|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all,\nscenarios.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "copy-icon",
							"description": "The icon to show in the default copy state. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "success-icon",
							"description": "The icon to show when the content is copied. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "error-icon",
							"description": "The icon to show when a copy error occurs. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-copy",
							"description": "Emitted when the data has been copied.",
						},
						bson.M{
							"name":        "sl-error",
							"description": "Emitted when the data could not be copied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "copyIcon",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "successIcon",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "errorIcon",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "tooltip",
								"type": "SlTooltip",
							},
							bson.M{
								"name": "isCopying",
								"type": "boolean",
							},
							bson.M{
								"name": "status",
								"type": "'rest' | 'success' | 'error'",
							},
							bson.M{
								"name":        "value",
								"description": "The text value to copy.",
								"type":        "string",
							},
							bson.M{
								"name":        "from",
								"description": "An id that references an element in the same document from which data will be copied. If both this and `value` are\npresent, this value will take precedence. By default, the target element's `textContent` will be copied. To copy an\nattribute, append the attribute name wrapped in square brackets, e.g. `from=\"elbson.A{value}\"`. To copy a property,\nappend a dot and the property name, e.g. `from=\"el.value\"`.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the copy button.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "copyLabel",
								"description": "A custom label to show in the tooltip.",
								"type":        "string",
							},
							bson.M{
								"name":        "successLabel",
								"description": "A custom label to show in the tooltip after copying.",
								"type":        "string",
							},
							bson.M{
								"name":        "errorLabel",
								"description": "A custom label to show in the tooltip when a copy error occurs.",
								"type":        "string",
							},
							bson.M{
								"name":        "feedbackDuration",
								"description": "The length of time to show feedback before restoring the default trigger.",
								"type":        "number",
							},
							bson.M{
								"name":        "tooltipPlacement",
								"description": "The preferred placement of the tooltip.",
								"type":        "'top' | 'right' | 'bottom' | 'left'",
							},
							bson.M{
								"name":        "hoist",
								"description": "Enable this option to prevent the tooltip from being clipped when the component is placed inside a container with\n`overflow: auto|hidden|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all,\nscenarios.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-copy",
								"description": "Emitted when the data has been copied.",
							},
							bson.M{
								"name":        "sl-error",
								"description": "Emitted when the data could not be copied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-details",
					"description": "Details show a brief summary and expand to show additional content.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the details opens.\n- **sl-after-show** - Emitted after the details opens and all animations are complete.\n- **sl-hide** - Emitted when the details closes.\n- **sl-after-hide** - Emitted after the details closes and all animations are complete.\n\n### **Methods:**\n - **show()** - Shows the details.\n- **hide()** - Hides the details\n\n### **Slots:**\n - _default_ - The details' main content.\n- **summary** - The details' summary. Alternatively, you can use the `summary` attribute.\n- **expand-icon** - Optional expand icon to use instead of the default. Works best with `<sl-icon>`.\n- **collapse-icon** - Optional collapse icon to use instead of the default. Works best with `<sl-icon>`.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **header** - The header that wraps both the summary and the expand/collapse icon.\n- **summary** - The container that wraps the summary.\n- **summary-icon** - The container that wraps the expand/collapse icons.\n- **content** - The details content.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the details is open. You can toggle this attribute to show and hide the details, or you\ncan use the `show()` and `hide()` methods and this attribute will reflect the details' open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "summary",
							"description": "The summary to show in the header. If you need to display HTML, use the `summary` slot instead.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the details so it can't be toggled.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The details' main content.",
						},
						bson.M{
							"name":        "summary",
							"description": "The details' summary. Alternatively, you can use the `summary` attribute.",
						},
						bson.M{
							"name":        "expand-icon",
							"description": "Optional expand icon to use instead of the default. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "collapse-icon",
							"description": "Optional collapse icon to use instead of the default. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the details opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the details opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the details closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the details closes and all animations are complete.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "details",
								"type": "HTMLDetailsElement",
							},
							bson.M{
								"name": "header",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "body",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "expandIconSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "detailsObserver",
								"type": "MutationObserver",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the details is open. You can toggle this attribute to show and hide the details, or you\ncan use the `show()` and `hide()` methods and this attribute will reflect the details' open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "summary",
								"description": "The summary to show in the header. If you need to display HTML, use the `summary` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the details so it can't be toggled.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the details opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the details opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the details closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the details closes and all animations are complete.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-dialog",
					"description": "Dialogs, sometimes called \"modals\", appear above the page and require the user's immediate attention.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the dialog opens.\n- **sl-after-show** - Emitted after the dialog opens and all animations are complete.\n- **sl-hide** - Emitted when the dialog closes.\n- **sl-after-hide** - Emitted after the dialog closes and all animations are complete.\n- **sl-initial-focus** - Emitted when the dialog opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.\n- **sl-request-close** - Emitted when the user attempts to close the dialog by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the dialog open. Avoid using this unless closing the dialog will result in destructive behavior such as data loss.\n\n### **Methods:**\n - **show()** - Shows the dialog.\n- **hide()** - Hides the dialog\n\n### **Slots:**\n - _default_ - The dialog's main content.\n- **label** - The dialog's label. Alternatively, you can use the `label` attribute.\n- **header-actions** - Optional actions to add to the header. Works best with `<sl-icon-button>`.\n- **footer** - The dialog's footer, usually one or more buttons representing various options.\n\n### **CSS Properties:**\n - **--width** - The preferred width of the dialog. Note that the dialog will shrink to accommodate smaller screens. _(default: undefined)_\n- **--header-spacing** - The amount of padding to use for the header. _(default: undefined)_\n- **--body-spacing** - The amount of padding to use for the body. _(default: undefined)_\n- **--footer-spacing** - The amount of padding to use for the footer. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **overlay** - The overlay that covers the screen behind the dialog.\n- **panel** - The dialog's panel (where the dialog and its content are rendered).\n- **header** - The dialog's header. This element wraps the title and header actions.\n- **header-actions** - Optional actions to add to the header. Works best with `<sl-icon-button>`.\n- **title** - The dialog's title.\n- **close-button** - The close button, an `<sl-icon-button>`.\n- **close-button__base** - The close button's exported `base` part.\n- **body** - The dialog's body.\n- **footer** - The dialog's footer.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the dialog is open. You can toggle this attribute to show and hide the dialog, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the dialog's open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The dialog's label as displayed in the header. You should always include a relevant label even when using\n`no-header`, as it is required for proper accessibility. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "no-header",
							"description": "Disables the header. This will also remove the default close button, so please ensure you provide an easy,\naccessible way for users to dismiss the dialog.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The dialog's main content.",
						},
						bson.M{
							"name":        "label",
							"description": "The dialog's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "header-actions",
							"description": "Optional actions to add to the header. Works best with `<sl-icon-button>`.",
						},
						bson.M{
							"name":        "footer",
							"description": "The dialog's footer, usually one or more buttons representing various options.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the dialog opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the dialog opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the dialog closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the dialog closes and all animations are complete.",
						},
						bson.M{
							"name":        "sl-initial-focus",
							"description": "Emitted when the dialog opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.",
						},
						bson.M{
							"name":        "sl-request-close",
							"type":        "bson.M{ source: 'close-button' | 'keyboard' | 'overlay' }",
							"description": "Emitted when the user attempts to close the dialog by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the dialog open. Avoid using this unless closing the dialog will result in destructive behavior such as data loss.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "modal",
								"description": "Exposes the internal modal utility that controls focus trapping. To temporarily disable focus trapping and allow third-party modals spawned from an active Shoelace modal, call `modal.activateExternal()` when the third-party modal opens. Upon closing, call `modal.deactivateExternal()` to restore Shoelace's focus trapping.",
							},
							bson.M{
								"name": "dialog",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "panel",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "overlay",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the dialog is open. You can toggle this attribute to show and hide the dialog, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the dialog's open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "The dialog's label as displayed in the header. You should always include a relevant label even when using\n`no-header`, as it is required for proper accessibility. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "noHeader",
								"description": "Disables the header. This will also remove the default close button, so please ensure you provide an easy,\naccessible way for users to dismiss the dialog.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the dialog opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the dialog opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the dialog closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the dialog closes and all animations are complete.",
							},
							bson.M{
								"name":        "sl-initial-focus",
								"description": "Emitted when the dialog opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.",
							},
							bson.M{
								"name":        "sl-request-close",
								"type":        "bson.M{ source: 'close-button' | 'keyboard' | 'overlay' }",
								"description": "Emitted when the user attempts to close the dialog by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the dialog open. Avoid using this unless closing the dialog will result in destructive behavior such as data loss.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-divider",
					"description": "Dividers are used to visually separate or group elements.\n---\n\n\n### **CSS Properties:**\n - **--color** - The color of the divider. _(default: undefined)_\n- **--width** - The width of the divider. _(default: undefined)_\n- **--spacing** - The spacing of the divider. _(default: undefined)_",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "vertical",
							"description": "Draws the divider in a vertical orientation.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "vertical",
								"description": "Draws the divider in a vertical orientation.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-drawer",
					"description": "Drawers slide in from a container to expose additional options and information.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the drawer opens.\n- **sl-after-show** - Emitted after the drawer opens and all animations are complete.\n- **sl-hide** - Emitted when the drawer closes.\n- **sl-after-hide** - Emitted after the drawer closes and all animations are complete.\n- **sl-initial-focus** - Emitted when the drawer opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.\n- **sl-request-close** - Emitted when the user attempts to close the drawer by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the drawer open. Avoid using this unless closing the drawer will result in destructive behavior such as data loss.\n\n### **Methods:**\n - **show()** - Shows the drawer.\n- **hide()** - Hides the drawer\n\n### **Slots:**\n - _default_ - The drawer's main content.\n- **label** - The drawer's label. Alternatively, you can use the `label` attribute.\n- **header-actions** - Optional actions to add to the header. Works best with `<sl-icon-button>`.\n- **footer** - The drawer's footer, usually one or more buttons representing various options.\n\n### **CSS Properties:**\n - **--size** - The preferred size of the drawer. This will be applied to the drawer's width or height depending on its `placement`. Note that the drawer will shrink to accommodate smaller screens. _(default: undefined)_\n- **--header-spacing** - The amount of padding to use for the header. _(default: undefined)_\n- **--body-spacing** - The amount of padding to use for the body. _(default: undefined)_\n- **--footer-spacing** - The amount of padding to use for the footer. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **overlay** - The overlay that covers the screen behind the drawer.\n- **panel** - The drawer's panel (where the drawer and its content are rendered).\n- **header** - The drawer's header. This element wraps the title and header actions.\n- **header-actions** - Optional actions to add to the header. Works best with `<sl-icon-button>`.\n- **title** - The drawer's title.\n- **close-button** - The close button, an `<sl-icon-button>`.\n- **close-button__base** - The close button's exported `base` part.\n- **body** - The drawer's body.\n- **footer** - The drawer's footer.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the drawer is open. You can toggle this attribute to show and hide the drawer, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the drawer's open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The drawer's label as displayed in the header. You should always include a relevant label even when using\n`no-header`, as it is required for proper accessibility. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "placement",
							"description": "The direction from which the drawer will open.",
							"value": bson.M{
								"type":    "'top' | 'end' | 'bottom' | 'start'",
								"default": "'end'",
							},
						},
						bson.M{
							"name":        "contained",
							"description": "By default, the drawer slides out of its containing block (usually the viewport). To make the drawer slide out of\nits parent element, set this attribute and add `position: relative` to the parent.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "no-header",
							"description": "Removes the header. This will also remove the default close button, so please ensure you provide an easy,\naccessible way for users to dismiss the drawer.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The drawer's main content.",
						},
						bson.M{
							"name":        "label",
							"description": "The drawer's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "header-actions",
							"description": "Optional actions to add to the header. Works best with `<sl-icon-button>`.",
						},
						bson.M{
							"name":        "footer",
							"description": "The drawer's footer, usually one or more buttons representing various options.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the drawer opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the drawer opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the drawer closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the drawer closes and all animations are complete.",
						},
						bson.M{
							"name":        "sl-initial-focus",
							"description": "Emitted when the drawer opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.",
						},
						bson.M{
							"name":        "sl-request-close",
							"type":        "bson.M{ source: 'close-button' | 'keyboard' | 'overlay' }",
							"description": "Emitted when the user attempts to close the drawer by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the drawer open. Avoid using this unless closing the drawer will result in destructive behavior such as data loss.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "modal",
								"description": "Exposes the internal modal utility that controls focus trapping. To temporarily disable focus trapping and allow third-party modals spawned from an active Shoelace modal, call `modal.activateExternal()` when the third-party modal opens. Upon closing, call `modal.deactivateExternal()` to restore Shoelace's focus trapping.",
							},
							bson.M{
								"name": "drawer",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "panel",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "overlay",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the drawer is open. You can toggle this attribute to show and hide the drawer, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the drawer's open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "The drawer's label as displayed in the header. You should always include a relevant label even when using\n`no-header`, as it is required for proper accessibility. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "placement",
								"description": "The direction from which the drawer will open.",
								"type":        "'top' | 'end' | 'bottom' | 'start'",
							},
							bson.M{
								"name":        "contained",
								"description": "By default, the drawer slides out of its containing block (usually the viewport). To make the drawer slide out of\nits parent element, set this attribute and add `position: relative` to the parent.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "noHeader",
								"description": "Removes the header. This will also remove the default close button, so please ensure you provide an easy,\naccessible way for users to dismiss the drawer.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the drawer opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the drawer opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the drawer closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the drawer closes and all animations are complete.",
							},
							bson.M{
								"name":        "sl-initial-focus",
								"description": "Emitted when the drawer opens and is ready to receive focus. Calling `event.preventDefault()` will prevent focusing and allow you to set it on a different element, such as an input.",
							},
							bson.M{
								"name":        "sl-request-close",
								"type":        "bson.M{ source: 'close-button' | 'keyboard' | 'overlay' }",
								"description": "Emitted when the user attempts to close the drawer by clicking the close button, clicking the overlay, or pressing escape. Calling `event.preventDefault()` will keep the drawer open. Avoid using this unless closing the drawer will result in destructive behavior such as data loss.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-dropdown",
					"description": "Dropdowns expose additional content that \"drops down\" in a panel.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the dropdown opens.\n- **sl-after-show** - Emitted after the dropdown opens and all animations are complete.\n- **sl-hide** - Emitted when the dropdown closes.\n- **sl-after-hide** - Emitted after the dropdown closes and all animations are complete.\n\n### **Methods:**\n - **show()** - Shows the dropdown panel.\n- **hide()** - Hides the dropdown panel\n- **reposition()** - Instructs the dropdown menu to reposition. Useful when the position or size of the trigger changes when the menu\nis activated.\n\n### **Slots:**\n - _default_ - The dropdown's main content.\n- **trigger** - The dropdown's trigger, usually a `<sl-button>` element.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper, an `<sl-popup>` element.\n- **base__popup** - The popup's exported `popup` part. Use this to target the tooltip's popup container.\n- **trigger** - The container that wraps the trigger.\n- **panel** - The panel that gets shown when the dropdown is open.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the dropdown is open. You can toggle this attribute to show and hide the dropdown, or you\ncan use the `show()` and `hide()` methods and this attribute will reflect the dropdown's open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "placement",
							"description": "The preferred placement of the dropdown panel. Note that the actual placement may vary as needed to keep the panel\ninside of the viewport.",
							"value": bson.M{
								"type":    "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
								"default": "'bottom-start'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the dropdown so the panel will not open.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "stay-open-on-select",
							"description": "By default, the dropdown is closed when an item is selected. This attribute will keep it open instead. Useful for\ndropdowns that allow for multiple interactions.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "distance",
							"description": "The distance in pixels from which to offset the panel away from its trigger.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "skidding",
							"description": "The distance in pixels from which to offset the panel along its trigger.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "hoist",
							"description": "Enable this option to prevent the panel from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "sync",
							"description": "Syncs the popup width or height to that of the trigger element.",
							"value": bson.M{
								"type":    "'width' | 'height' | 'both' | undefined",
								"default": "undefined",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The dropdown's main content.",
						},
						bson.M{
							"name":        "trigger",
							"description": "The dropdown's trigger, usually a `<sl-button>` element.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the dropdown opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the dropdown opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the dropdown closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the dropdown closes and all animations are complete.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "popup",
								"type": "SlPopup",
							},
							bson.M{
								"name": "trigger",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "panel",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the dropdown is open. You can toggle this attribute to show and hide the dropdown, or you\ncan use the `show()` and `hide()` methods and this attribute will reflect the dropdown's open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "placement",
								"description": "The preferred placement of the dropdown panel. Note that the actual placement may vary as needed to keep the panel\ninside of the viewport.",
								"type":        "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the dropdown so the panel will not open.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "stayOpenOnSelect",
								"description": "By default, the dropdown is closed when an item is selected. This attribute will keep it open instead. Useful for\ndropdowns that allow for multiple interactions.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "containingElement",
								"description": "The dropdown will close when the user interacts outside of this element (e.g. clicking). Useful for composing other\ncomponents that use a dropdown internally.",
								"type":        "HTMLElement | undefined",
							},
							bson.M{
								"name":        "distance",
								"description": "The distance in pixels from which to offset the panel away from its trigger.",
								"type":        "number",
							},
							bson.M{
								"name":        "skidding",
								"description": "The distance in pixels from which to offset the panel along its trigger.",
								"type":        "number",
							},
							bson.M{
								"name":        "hoist",
								"description": "Enable this option to prevent the panel from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "sync",
								"description": "Syncs the popup width or height to that of the trigger element.",
								"type":        "'width' | 'height' | 'both' | undefined",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the dropdown opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the dropdown opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the dropdown closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the dropdown closes and all animations are complete.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-format-bytes",
					"description": "Formats a number as a human readable bytes value.\n---\n",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The number to format in bytes.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "unit",
							"description": "The type of unit to display.",
							"value": bson.M{
								"type":    "'byte' | 'bit'",
								"default": "'byte'",
							},
						},
						bson.M{
							"name":        "display",
							"description": "Determines how to display the result, e.g. \"100 bytes\", \"100 b\", or \"100b\".",
							"value": bson.M{
								"type":    "'long' | 'short' | 'narrow'",
								"default": "'short'",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "value",
								"description": "The number to format in bytes.",
								"type":        "number",
							},
							bson.M{
								"name":        "unit",
								"description": "The type of unit to display.",
								"type":        "'byte' | 'bit'",
							},
							bson.M{
								"name":        "display",
								"description": "Determines how to display the result, e.g. \"100 bytes\", \"100 b\", or \"100b\".",
								"type":        "'long' | 'short' | 'narrow'",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-format-date",
					"description": "Formats a date/time using the specified locale and options.\n---\n",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "date",
							"description": "The date/time to format. If not set, the current date and time will be used. When passing a string, it's strongly\nrecommended to use the ISO 8601 format to ensure timezones are handled correctly. To convert a date to this format\nin JavaScript, use bson.A{`date.toISOString()`}(https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString).",
							"value": bson.M{
								"type":    "Date | string",
								"default": "new Date()",
							},
						},
						bson.M{
							"name":        "weekday",
							"description": "The format for displaying the weekday.",
							"value": bson.M{
								"type": "'narrow' | 'short' | 'long'",
							},
						},
						bson.M{
							"name":        "era",
							"description": "The format for displaying the era.",
							"value": bson.M{
								"type": "'narrow' | 'short' | 'long'",
							},
						},
						bson.M{
							"name":        "year",
							"description": "The format for displaying the year.",
							"value": bson.M{
								"type": "'numeric' | '2-digit'",
							},
						},
						bson.M{
							"name":        "month",
							"description": "The format for displaying the month.",
							"value": bson.M{
								"type": "'numeric' | '2-digit' | 'narrow' | 'short' | 'long'",
							},
						},
						bson.M{
							"name":        "day",
							"description": "The format for displaying the day.",
							"value": bson.M{
								"type": "'numeric' | '2-digit'",
							},
						},
						bson.M{
							"name":        "hour",
							"description": "The format for displaying the hour.",
							"value": bson.M{
								"type": "'numeric' | '2-digit'",
							},
						},
						bson.M{
							"name":        "minute",
							"description": "The format for displaying the minute.",
							"value": bson.M{
								"type": "'numeric' | '2-digit'",
							},
						},
						bson.M{
							"name":        "second",
							"description": "The format for displaying the second.",
							"value": bson.M{
								"type": "'numeric' | '2-digit'",
							},
						},
						bson.M{
							"name":        "time-zone-name",
							"description": "The format for displaying the time.",
							"value": bson.M{
								"type": "'short' | 'long'",
							},
						},
						bson.M{
							"name":        "time-zone",
							"description": "The time zone to express the time in.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "hour-format",
							"description": "The format for displaying the hour.",
							"value": bson.M{
								"type":    "'auto' | '12' | '24'",
								"default": "'auto'",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "date",
								"description": "The date/time to format. If not set, the current date and time will be used. When passing a string, it's strongly\nrecommended to use the ISO 8601 format to ensure timezones are handled correctly. To convert a date to this format\nin JavaScript, use bson.A{`date.toISOString()`}(https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString).",
								"type":        "Date | string",
							},
							bson.M{
								"name":        "weekday",
								"description": "The format for displaying the weekday.",
								"type":        "'narrow' | 'short' | 'long'",
							},
							bson.M{
								"name":        "era",
								"description": "The format for displaying the era.",
								"type":        "'narrow' | 'short' | 'long'",
							},
							bson.M{
								"name":        "year",
								"description": "The format for displaying the year.",
								"type":        "'numeric' | '2-digit'",
							},
							bson.M{
								"name":        "month",
								"description": "The format for displaying the month.",
								"type":        "'numeric' | '2-digit' | 'narrow' | 'short' | 'long'",
							},
							bson.M{
								"name":        "day",
								"description": "The format for displaying the day.",
								"type":        "'numeric' | '2-digit'",
							},
							bson.M{
								"name":        "hour",
								"description": "The format for displaying the hour.",
								"type":        "'numeric' | '2-digit'",
							},
							bson.M{
								"name":        "minute",
								"description": "The format for displaying the minute.",
								"type":        "'numeric' | '2-digit'",
							},
							bson.M{
								"name":        "second",
								"description": "The format for displaying the second.",
								"type":        "'numeric' | '2-digit'",
							},
							bson.M{
								"name":        "timeZoneName",
								"description": "The format for displaying the time.",
								"type":        "'short' | 'long'",
							},
							bson.M{
								"name":        "timeZone",
								"description": "The time zone to express the time in.",
								"type":        "string",
							},
							bson.M{
								"name":        "hourFormat",
								"description": "The format for displaying the hour.",
								"type":        "'auto' | '12' | '24'",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-format-number",
					"description": "Formats a number using the specified locale and options.\n---\n",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The number to format.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "type",
							"description": "The formatting style to use.",
							"value": bson.M{
								"type":    "'currency' | 'decimal' | 'percent'",
								"default": "'decimal'",
							},
						},
						bson.M{
							"name":        "no-grouping",
							"description": "Turns off grouping separators.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "currency",
							"description": "The bson.A{ISO 4217}(https://en.wikipedia.org/wiki/ISO_4217) currency code to use when formatting.",
							"value": bson.M{
								"type":    "string",
								"default": "'USD'",
							},
						},
						bson.M{
							"name":        "currency-display",
							"description": "How to display the currency.",
							"value": bson.M{
								"type":    "'symbol' | 'narrowSymbol' | 'code' | 'name'",
								"default": "'symbol'",
							},
						},
						bson.M{
							"name":        "minimum-integer-digits",
							"description": "The minimum number of integer digits to use. Possible values are 1-21.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "minimum-fraction-digits",
							"description": "The minimum number of fraction digits to use. Possible values are 0-20.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "maximum-fraction-digits",
							"description": "The maximum number of fraction digits to use. Possible values are 0-0.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "minimum-significant-digits",
							"description": "The minimum number of significant digits to use. Possible values are 1-21.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "maximum-significant-digits",
							"description": "The maximum number of significant digits to use,. Possible values are 1-21.",
							"value": bson.M{
								"type": "number",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "value",
								"description": "The number to format.",
								"type":        "number",
							},
							bson.M{
								"name":        "type",
								"description": "The formatting style to use.",
								"type":        "'currency' | 'decimal' | 'percent'",
							},
							bson.M{
								"name":        "noGrouping",
								"description": "Turns off grouping separators.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "currency",
								"description": "The bson.A{ISO 4217}(https://en.wikipedia.org/wiki/ISO_4217) currency code to use when formatting.",
								"type":        "string",
							},
							bson.M{
								"name":        "currencyDisplay",
								"description": "How to display the currency.",
								"type":        "'symbol' | 'narrowSymbol' | 'code' | 'name'",
							},
							bson.M{
								"name":        "minimumIntegerDigits",
								"description": "The minimum number of integer digits to use. Possible values are 1-21.",
								"type":        "number",
							},
							bson.M{
								"name":        "minimumFractionDigits",
								"description": "The minimum number of fraction digits to use. Possible values are 0-20.",
								"type":        "number",
							},
							bson.M{
								"name":        "maximumFractionDigits",
								"description": "The maximum number of fraction digits to use. Possible values are 0-0.",
								"type":        "number",
							},
							bson.M{
								"name":        "minimumSignificantDigits",
								"description": "The minimum number of significant digits to use. Possible values are 1-21.",
								"type":        "number",
							},
							bson.M{
								"name":        "maximumSignificantDigits",
								"description": "The maximum number of significant digits to use,. Possible values are 1-21.",
								"type":        "number",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-icon",
					"description": "Icons are symbols that can be used to represent various options within an application.\n---\n\n\n### **Events:**\n - **sl-load** - Emitted when the icon has loaded. When using `spriteSheet: true` this will not emit.\n- **sl-error** - Emitted when the icon fails to load due to an error. When using `spriteSheet: true` this will not emit.\n\n### **CSS Parts:**\n - **svg** - The internal SVG element.\n- **use** - The <use> element generated when using `spriteSheet: true`",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "name",
							"description": "The name of the icon to draw. Available names depend on the icon library being used.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "src",
							"description": "An external URL of an SVG file. Be sure you trust the content you are including, as it will be executed as code and\ncan result in XSS attacks.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "label",
							"description": "An alternate description to use for assistive devices. If omitted, the icon will be considered presentational and\nignored by assistive devices.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "library",
							"description": "The name of a registered custom icon library.",
							"value": bson.M{
								"type":    "string",
								"default": "'default'",
							},
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-load",
							"description": "Emitted when the icon has loaded. When using `spriteSheet: true` this will not emit.",
						},
						bson.M{
							"name":        "sl-error",
							"description": "Emitted when the icon fails to load due to an error. When using `spriteSheet: true` this will not emit.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "name",
								"description": "The name of the icon to draw. Available names depend on the icon library being used.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "src",
								"description": "An external URL of an SVG file. Be sure you trust the content you are including, as it will be executed as code and\ncan result in XSS attacks.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "label",
								"description": "An alternate description to use for assistive devices. If omitted, the icon will be considered presentational and\nignored by assistive devices.",
								"type":        "string",
							},
							bson.M{
								"name":        "library",
								"description": "The name of a registered custom icon library.",
								"type":        "string",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-load",
								"description": "Emitted when the icon has loaded. When using `spriteSheet: true` this will not emit.",
							},
							bson.M{
								"name":        "sl-error",
								"description": "Emitted when the icon fails to load due to an error. When using `spriteSheet: true` this will not emit.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-icon-button",
					"description": "Icons buttons are simple, icon-only buttons that can be used for actions and in toolbars.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the icon button loses focus.\n- **sl-focus** - Emitted when the icon button gains focus.\n\n### **Methods:**\n - **click()** - Simulates a click on the icon button.\n- **focus(options: _FocusOptions_)** - Sets focus on the icon button.\n- **blur()** - Removes focus from the icon button.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "name",
							"description": "The name of the icon to draw. Available names depend on the icon library being used.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "library",
							"description": "The name of a registered custom icon library.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "src",
							"description": "An external URL of an SVG file. Be sure you trust the content you are including, as it will be executed as code and\ncan result in XSS attacks.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "href",
							"description": "When set, the underlying button will be rendered as an `<a>` with this `href` instead of a `<button>`.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "target",
							"description": "Tells the browser where to open the link. Only used when `href` is set.",
							"value": bson.M{
								"type": "'_blank' | '_parent' | '_self' | '_top' | undefined",
							},
						},
						bson.M{
							"name":        "download",
							"description": "Tells the browser to download the linked file as this filename. Only used when `href` is set.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "label",
							"description": "A description that gets read by assistive devices. For optimal accessibility, you should always include a label\nthat describes what the icon button does.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the icon button loses focus.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the icon button gains focus.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "button",
								"type": "HTMLButtonElement | HTMLLinkElement",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the icon to draw. Available names depend on the icon library being used.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "library",
								"description": "The name of a registered custom icon library.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "src",
								"description": "An external URL of an SVG file. Be sure you trust the content you are including, as it will be executed as code and\ncan result in XSS attacks.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "href",
								"description": "When set, the underlying button will be rendered as an `<a>` with this `href` instead of a `<button>`.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "target",
								"description": "Tells the browser where to open the link. Only used when `href` is set.",
								"type":        "'_blank' | '_parent' | '_self' | '_top' | undefined",
							},
							bson.M{
								"name":        "download",
								"description": "Tells the browser to download the linked file as this filename. Only used when `href` is set.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "label",
								"description": "A description that gets read by assistive devices. For optimal accessibility, you should always include a label\nthat describes what the icon button does.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the button.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the icon button loses focus.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the icon button gains focus.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-image-comparer",
					"description": "Compare visual differences between similar photos with a sliding panel.\n---\n\n\n### **Events:**\n - **sl-change** - Emitted when the position changes.\n\n### **Slots:**\n - **before** - The before image, an `<img>` or `<svg>` element.\n- **after** - The after image, an `<img>` or `<svg>` element.\n- **handle** - The icon used inside the handle.\n\n### **CSS Properties:**\n - **--divider-width** - The width of the dividing line. _(default: undefined)_\n- **--handle-size** - The size of the compare handle. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **before** - The container that wraps the before image.\n- **after** - The container that wraps the after image.\n- **divider** - The divider that separates the images.\n- **handle** - The handle that the user drags to expose the after image.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "position",
							"description": "The position of the divider as a percentage.",
							"value": bson.M{
								"type":    "number",
								"default": "50",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "before",
							"description": "The before image, an `<img>` or `<svg>` element.",
						},
						bson.M{
							"name":        "after",
							"description": "The after image, an `<img>` or `<svg>` element.",
						},
						bson.M{
							"name":        "handle",
							"description": "The icon used inside the handle.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the position changes.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "base",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "handle",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "position",
								"description": "The position of the divider as a percentage.",
								"type":        "number",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the position changes.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-include",
					"description": "Includes give you the power to embed external HTML files into the page.\n---\n\n\n### **Events:**\n - **sl-load** - Emitted when the included file is loaded.\n- **sl-error** - Emitted when the included file fails to load due to an error.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "src",
							"description": "The location of the HTML file to include. Be sure you trust the content you are including as it will be executed as\ncode and can result in XSS attacks.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "mode",
							"description": "The fetch mode to use.",
							"value": bson.M{
								"type":    "'cors' | 'no-cors' | 'same-origin'",
								"default": "'cors'",
							},
						},
						bson.M{
							"name":        "allow-scripts",
							"description": "Allows included scripts to be executed. Be sure you trust the content you are including as it will be executed as\ncode and can result in XSS attacks.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-load",
							"description": "Emitted when the included file is loaded.",
						},
						bson.M{
							"name":        "sl-error",
							"type":        "bson.M{ status: number }",
							"description": "Emitted when the included file fails to load due to an error.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "src",
								"description": "The location of the HTML file to include. Be sure you trust the content you are including as it will be executed as\ncode and can result in XSS attacks.",
								"type":        "string",
							},
							bson.M{
								"name":        "mode",
								"description": "The fetch mode to use.",
								"type":        "'cors' | 'no-cors' | 'same-origin'",
							},
							bson.M{
								"name":        "allowScripts",
								"description": "Allows included scripts to be executed. Be sure you trust the content you are including as it will be executed as\ncode and can result in XSS attacks.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-load",
								"description": "Emitted when the included file is loaded.",
							},
							bson.M{
								"name":        "sl-error",
								"type":        "bson.M{ status: number }",
								"description": "Emitted when the included file fails to load due to an error.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-input",
					"description": "Inputs collect data from the user.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the control loses focus.\n- **sl-change** - Emitted when an alteration to the control's value is committed by the user.\n- **sl-clear** - Emitted when the clear button is activated.\n- **sl-focus** - Emitted when the control gains focus.\n- **sl-input** - Emitted when the control receives input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the input.\n- **blur()** - Removes focus from the input.\n- **select()** - Selects all the text in the input.\n- **setSelectionRange(selectionStart: _number_, selectionEnd: _number_, selectionDirection: _'forward' | 'backward' | 'none'_)** - Sets the start and end positions of the text selection (0-based).\n- **setRangeText(replacement: _string_, start: _number_, end: _number_, selectMode: _'select' | 'start' | 'end' | 'preserve'_)** - Replaces a range of text with a new string.\n- **showPicker()** - Displays the browser picker for an input element (only works if the browser supports it for the input type).\n- **stepUp()** - Increments the value of a numeric input type by the value of the step attribute.\n- **stepDown()** - Decrements the value of a numeric input type by the value of the step attribute.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - **label** - The input's label. Alternatively, you can use the `label` attribute.\n- **prefix** - Used to prepend a presentational icon or similar element to the input.\n- **suffix** - Used to append a presentational icon or similar element to the input.\n- **clear-icon** - An icon to use in lieu of the default clear icon.\n- **show-password-icon** - An icon to use in lieu of the default show password icon.\n- **hide-password-icon** - An icon to use in lieu of the default hide password icon.\n- **help-text** - Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Parts:**\n - **form-control** - The form control that wraps the label, input, and help text.\n- **form-control-label** - The label's wrapper.\n- **form-control-input** - The input's wrapper.\n- **form-control-help-text** - The help text's wrapper.\n- **base** - The component's base wrapper.\n- **input** - The internal `<input>` control.\n- **prefix** - The container that wraps the prefix.\n- **clear-button** - The clear button.\n- **password-toggle-button** - The password toggle button.\n- **suffix** - The container that wraps the suffix.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "type",
							"description": "The type of input. Works the same as a native `<input>` element, but only a subset of types are supported. Defaults\nto `text`.",
							"value": bson.M{
								"type":    "| 'date'\n    | 'datetime-local'\n    | 'email'\n    | 'number'\n    | 'password'\n    | 'search'\n    | 'tel'\n    | 'text'\n    | 'time'\n    | 'url'",
								"default": "'text'",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the input, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the input, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The input's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "filled",
							"description": "Draws a filled input.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style input with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The input's label. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The input's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "clearable",
							"description": "Adds a clear button when the input is not empty.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the input.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "placeholder",
							"description": "Placeholder text to show as a hint when the input is empty.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "readonly",
							"description": "Makes the input readonly.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "password-toggle",
							"description": "Adds a button to toggle the password's visibility. Only applies to password types.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "password-visible",
							"description": "Determines whether or not the password is currently visible. Only applies to password input types.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "no-spin-buttons",
							"description": "Hides the browser's built-in increment/decrement spin buttons for number inputs.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Makes the input a required field.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pattern",
							"description": "A regular expression pattern to validate input against.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "minlength",
							"description": "The minimum length of input that will be considered valid.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "maxlength",
							"description": "The maximum length of input that will be considered valid.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "min",
							"description": "The input's minimum value. Only applies to date and number input types.",
							"value": bson.M{
								"type": "number | string",
							},
						},
						bson.M{
							"name":        "max",
							"description": "The input's maximum value. Only applies to date and number input types.",
							"value": bson.M{
								"type": "number | string",
							},
						},
						bson.M{
							"name":        "step",
							"description": "Specifies the granularity that the value must adhere to, or the special value `any` which means no stepping is\nimplied, allowing any numeric value. Only applies to date and number input types.",
							"value": bson.M{
								"type": "number | 'any'",
							},
						},
						bson.M{
							"name":        "autocapitalize",
							"description": "Controls whether and how text input is automatically capitalized as it is entered by the user.",
							"value": bson.M{
								"type": "'off' | 'none' | 'on' | 'sentences' | 'words' | 'characters'",
							},
						},
						bson.M{
							"name":        "autocorrect",
							"description": "Indicates whether the browser's autocorrect feature is on or off.",
							"value": bson.M{
								"type": "'off' | 'on'",
							},
						},
						bson.M{
							"name":        "autocomplete",
							"description": "Specifies what permission the browser has to provide assistance in filling out form field values. Refer to\nbson.A{this page on MDN}(https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) for available values.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "autofocus",
							"description": "Indicates that the input should receive focus on page load.",
							"value": bson.M{
								"type": "boolean",
							},
						},
						bson.M{
							"name":        "enterkeyhint",
							"description": "Used to customize the label or icon of the Enter key on virtual keyboards.",
							"value": bson.M{
								"type": "'enter' | 'done' | 'go' | 'next' | 'previous' | 'search' | 'send'",
							},
						},
						bson.M{
							"name":        "spellcheck",
							"description": "Enables spell checking on the input.",
							"value": bson.M{
								"type":    "boolean",
								"default": "true",
							},
						},
						bson.M{
							"name":        "inputmode",
							"description": "Tells the browser what type of data will be entered by the user, allowing it to display the appropriate virtual\nkeyboard on supportive devices.",
							"value": bson.M{
								"type": "'none' | 'text' | 'decimal' | 'numeric' | 'tel' | 'search' | 'email' | 'url'",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "label",
							"description": "The input's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "prefix",
							"description": "Used to prepend a presentational icon or similar element to the input.",
						},
						bson.M{
							"name":        "suffix",
							"description": "Used to append a presentational icon or similar element to the input.",
						},
						bson.M{
							"name":        "clear-icon",
							"description": "An icon to use in lieu of the default clear icon.",
						},
						bson.M{
							"name":        "show-password-icon",
							"description": "An icon to use in lieu of the default show password icon.",
						},
						bson.M{
							"name":        "hide-password-icon",
							"description": "An icon to use in lieu of the default hide password icon.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when an alteration to the control's value is committed by the user.",
						},
						bson.M{
							"name":        "sl-clear",
							"description": "Emitted when the clear button is activated.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the control receives input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "type",
								"description": "The type of input. Works the same as a native `<input>` element, but only a subset of types are supported. Defaults\nto `text`.",
								"type":        "| 'date'\n    | 'datetime-local'\n    | 'email'\n    | 'number'\n    | 'password'\n    | 'search'\n    | 'tel'\n    | 'text'\n    | 'time'\n    | 'url'",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the input, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the input, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "defaultValue",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The input's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "filled",
								"description": "Draws a filled input.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style input with rounded edges.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "The input's label. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "helpText",
								"description": "The input's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "clearable",
								"description": "Adds a clear button when the input is not empty.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the input.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "placeholder",
								"description": "Placeholder text to show as a hint when the input is empty.",
								"type":        "string",
							},
							bson.M{
								"name":        "readonly",
								"description": "Makes the input readonly.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "passwordToggle",
								"description": "Adds a button to toggle the password's visibility. Only applies to password types.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "passwordVisible",
								"description": "Determines whether or not the password is currently visible. Only applies to password input types.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "noSpinButtons",
								"description": "Hides the browser's built-in increment/decrement spin buttons for number inputs.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Makes the input a required field.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pattern",
								"description": "A regular expression pattern to validate input against.",
								"type":        "string",
							},
							bson.M{
								"name":        "minlength",
								"description": "The minimum length of input that will be considered valid.",
								"type":        "number",
							},
							bson.M{
								"name":        "maxlength",
								"description": "The maximum length of input that will be considered valid.",
								"type":        "number",
							},
							bson.M{
								"name":        "min",
								"description": "The input's minimum value. Only applies to date and number input types.",
								"type":        "number | string",
							},
							bson.M{
								"name":        "max",
								"description": "The input's maximum value. Only applies to date and number input types.",
								"type":        "number | string",
							},
							bson.M{
								"name":        "step",
								"description": "Specifies the granularity that the value must adhere to, or the special value `any` which means no stepping is\nimplied, allowing any numeric value. Only applies to date and number input types.",
								"type":        "number | 'any'",
							},
							bson.M{
								"name":        "autocapitalize",
								"description": "Controls whether and how text input is automatically capitalized as it is entered by the user.",
								"type":        "'off' | 'none' | 'on' | 'sentences' | 'words' | 'characters'",
							},
							bson.M{
								"name":        "autocorrect",
								"description": "Indicates whether the browser's autocorrect feature is on or off.",
								"type":        "'off' | 'on'",
							},
							bson.M{
								"name":        "autocomplete",
								"description": "Specifies what permission the browser has to provide assistance in filling out form field values. Refer to\nbson.A{this page on MDN}(https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) for available values.",
								"type":        "string",
							},
							bson.M{
								"name":        "autofocus",
								"description": "Indicates that the input should receive focus on page load.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "enterkeyhint",
								"description": "Used to customize the label or icon of the Enter key on virtual keyboards.",
								"type":        "'enter' | 'done' | 'go' | 'next' | 'previous' | 'search' | 'send'",
							},
							bson.M{
								"name":        "spellcheck",
								"description": "Enables spell checking on the input.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "inputmode",
								"description": "Tells the browser what type of data will be entered by the user, allowing it to display the appropriate virtual\nkeyboard on supportive devices.",
								"type":        "'none' | 'text' | 'decimal' | 'numeric' | 'tel' | 'search' | 'email' | 'url'",
							},
							bson.M{
								"name":        "valueAsDate",
								"description": "Gets or sets the current value as a `Date` object. Returns `null` if the value can't be converted. This will use the native `<input type=\"bson.M{bson.M{type}}\">` implementation and may result in an error.",
							},
							bson.M{
								"name":        "valueAsNumber",
								"description": "Gets or sets the current value as a number. Returns `NaN` if the value can't be converted.",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when an alteration to the control's value is committed by the user.",
							},
							bson.M{
								"name":        "sl-clear",
								"description": "Emitted when the clear button is activated.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the control receives input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-menu",
					"description": "Menus provide a list of options for the user to choose from.\n---\n\n\n### **Events:**\n - **sl-select** - Emitted when a menu item is selected.\n\n### **Slots:**\n - _default_ - The menu's content, including menu items, menu labels, and dividers.",
					"doc-url":     "",
					"attributes":  bson.A{},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The menu's content, including menu items, menu labels, and dividers.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-select",
							"type":        "bson.M{ item: SlMenuItem }",
							"description": "Emitted when a menu item is selected.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-select",
								"type":        "bson.M{ item: SlMenuItem }",
								"description": "Emitted when a menu item is selected.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-menu-label",
					"description": "Menu labels are used to describe a group of menu items.\n---\n\n\n### **Slots:**\n - _default_ - The menu label's content.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes":  bson.A{},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The menu label's content.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{},
						"events":     bson.A{},
					},
				},
				bson.M{
					"name":        "sl-mutation-observer",
					"description": "The Mutation Observer component offers a thin, declarative interface to the bson.A{`MutationObserver API`}(https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver).\n---\n\n\n### **Events:**\n - **sl-mutation** - Emitted when a mutation occurs.\n\n### **Slots:**\n - _default_ - The content to watch for mutations.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "attr",
							"description": "Watches for changes to attributes. To watch only specific attributes, separate them by a space, e.g.\n`attr=\"class id title\"`. To watch all attributes, use `*`.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "attr-old-value",
							"description": "Indicates whether or not the attribute's previous value should be recorded when monitoring changes.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "char-data",
							"description": "Watches for changes to the character data contained within the node.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "char-data-old-value",
							"description": "Indicates whether or not the previous value of the node's text should be recorded.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "child-list",
							"description": "Watches for the addition or removal of new child nodes.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the observer.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The content to watch for mutations.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-mutation",
							"type":        "bson.M{ mutationList: MutationRecordbson.A{} }",
							"description": "Emitted when a mutation occurs.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "attr",
								"description": "Watches for changes to attributes. To watch only specific attributes, separate them by a space, e.g.\n`attr=\"class id title\"`. To watch all attributes, use `*`.",
								"type":        "string",
							},
							bson.M{
								"name":        "attrOldValue",
								"description": "Indicates whether or not the attribute's previous value should be recorded when monitoring changes.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "charData",
								"description": "Watches for changes to the character data contained within the node.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "charDataOldValue",
								"description": "Indicates whether or not the previous value of the node's text should be recorded.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "childList",
								"description": "Watches for the addition or removal of new child nodes.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the observer.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-mutation",
								"type":        "bson.M{ mutationList: MutationRecordbson.A{} }",
								"description": "Emitted when a mutation occurs.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-option",
					"description": "Options define the selectable items within various form controls such as bson.A{select}(/components/select).\n---\n\n\n### **Methods:**\n - **getTextLabel()** - Returns a plain text label based on the option's content.\n\n### **Slots:**\n - _default_ - The option's label.\n- **prefix** - Used to prepend an icon or similar element to the menu item.\n- **suffix** - Used to append an icon or similar element to the menu item.\n\n### **CSS Parts:**\n - **checked-icon** - The checked icon, an `<sl-icon>` element.\n- **base** - The component's base wrapper.\n- **label** - The option's label.\n- **prefix** - The container that wraps the prefix.\n- **suffix** - The container that wraps the suffix.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The option's value. When selected, the containing form control will receive this value. The value must be unique\nfrom other options in the same group. Values may not contain spaces, as spaces are used as delimiters when listing\nmultiple values.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Draws the option in a disabled state, preventing selection.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The option's label.",
						},
						bson.M{
							"name":        "prefix",
							"description": "Used to prepend an icon or similar element to the menu item.",
						},
						bson.M{
							"name":        "suffix",
							"description": "Used to append an icon or similar element to the menu item.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "current",
								"type": "boolean",
							},
							bson.M{
								"name": "selected",
								"type": "boolean",
							},
							bson.M{
								"name": "hasHover",
								"type": "boolean",
							},
							bson.M{
								"name":        "value",
								"description": "The option's value. When selected, the containing form control will receive this value. The value must be unique\nfrom other options in the same group. Values may not contain spaces, as spaces are used as delimiters when listing\nmultiple values.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Draws the option in a disabled state, preventing selection.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-popup",
					"description": "Popup is a utility that lets you declaratively anchor \"popup\" containers to another element.\n---\n\n\n### **Events:**\n - **sl-reposition** - Emitted when the popup is repositioned. This event can fire a lot, so avoid putting expensive operations in your listener or consider debouncing it.\n\n### **Methods:**\n - **reposition()** - Forces the popup to recalculate and reposition itself.\n\n### **Slots:**\n - _default_ - The popup's content.\n- **anchor** - The element the popup will be anchored to. If the anchor lives outside of the popup, you can use the `anchor` attribute or property instead.\n\n### **CSS Properties:**\n - **--arrow-size** - The size of the arrow. Note that an arrow won't be shown unless the `arrow` attribute is used. _(default: 6px)_\n- **--arrow-color** - The color of the arrow. _(default: var(--sl-color-neutral-0))_\n- **--auto-size-available-width** - A read-only custom property that determines the amount of width the popup can be before overflowing. Useful for positioning child elements that need to overflow. This property is only available when using `auto-size`. _(default: undefined)_\n- **--auto-size-available-height** - A read-only custom property that determines the amount of height the popup can be before overflowing. Useful for positioning child elements that need to overflow. This property is only available when using `auto-size`. _(default: undefined)_\n\n### **CSS Parts:**\n - **arrow** - The arrow's container. Avoid setting `top|bottom|left|right` properties, as these values are assigned dynamically as the popup moves. This is most useful for applying a background color to match the popup, and maybe a border or box shadow.\n- **popup** - The popup's container. Useful for setting a background color, box shadow, etc.\n- **hover-bridge** - The hover bridge element. Only available when the `hover-bridge` option is enabled.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "anchor",
							"description": "The element the popup will be anchored to. If the anchor lives outside of the popup, you can provide the anchor\nelement `id`, a DOM element reference, or a `VirtualElement`. If the anchor lives inside the popup, use the\n`anchor` slot instead.",
							"value": bson.M{
								"type": "Element | string | VirtualElement",
							},
						},
						bson.M{
							"name":        "active",
							"description": "Activates the positioning logic and shows the popup. When this attribute is removed, the positioning logic is torn\ndown and the popup will be hidden.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "placement",
							"description": "The preferred placement of the popup. Note that the actual placement will vary as configured to keep the\npanel inside of the viewport.",
							"value": bson.M{
								"type":    "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
								"default": "'top'",
							},
						},
						bson.M{
							"name":        "strategy",
							"description": "Determines how the popup is positioned. The `absolute` strategy works well in most cases, but if overflow is\nclipped, using a `fixed` position strategy can often workaround it.",
							"value": bson.M{
								"type":    "'absolute' | 'fixed'",
								"default": "'absolute'",
							},
						},
						bson.M{
							"name":        "distance",
							"description": "The distance in pixels from which to offset the panel away from its anchor.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "skidding",
							"description": "The distance in pixels from which to offset the panel along its anchor.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "arrow",
							"description": "Attaches an arrow to the popup. The arrow's size and color can be customized using the `--arrow-size` and\n`--arrow-color` custom properties. For additional customizations, you can also target the arrow using\n`::part(arrow)` in your stylesheet.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "arrow-placement",
							"description": "The placement of the arrow. The default is `anchor`, which will align the arrow as close to the center of the\nanchor as possible, considering available space and `arrow-padding`. A value of `start`, `end`, or `center` will\nalign the arrow to the start, end, or center of the popover instead.",
							"value": bson.M{
								"type":    "'start' | 'end' | 'center' | 'anchor'",
								"default": "'anchor'",
							},
						},
						bson.M{
							"name":        "arrow-padding",
							"description": "The amount of padding between the arrow and the edges of the popup. If the popup has a border-radius, for example,\nthis will prevent it from overflowing the corners.",
							"value": bson.M{
								"type":    "number",
								"default": "10",
							},
						},
						bson.M{
							"name":        "flip",
							"description": "When set, placement of the popup will flip to the opposite site to keep it in view. You can use\n`flipFallbackPlacements` to further configure how the fallback placement is determined.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "flip-fallback-placements",
							"description": "If the preferred placement doesn't fit, popup will be tested in these fallback placements until one fits. Must be a\nstring of any number of placements separated by a space, e.g. \"top bottom left\". If no placement fits, the flip\nfallback strategy will be used instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "flip-fallback-strategy",
							"description": "When neither the preferred placement nor the fallback placements fit, this value will be used to determine whether\nthe popup should be positioned using the best available fit based on available space or as it was initially\npreferred.",
							"value": bson.M{
								"type":    "'best-fit' | 'initial'",
								"default": "'best-fit'",
							},
						},
						bson.M{
							"name":        "flipBoundary",
							"description": "The flip boundary describes clipping element(s) that overflow will be checked relative to when flipping. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
							"value": bson.M{
								"type": "Element | Elementbson.A{}",
							},
						},
						bson.M{
							"name":        "flip-padding",
							"description": "The amount of padding, in pixels, to exceed before the flip behavior will occur.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "shift",
							"description": "Moves the popup along the axis to keep it in view when clipped.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "shiftBoundary",
							"description": "The shift boundary describes clipping element(s) that overflow will be checked relative to when shifting. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
							"value": bson.M{
								"type": "Element | Elementbson.A{}",
							},
						},
						bson.M{
							"name":        "shift-padding",
							"description": "The amount of padding, in pixels, to exceed before the shift behavior will occur.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "auto-size",
							"description": "When set, this will cause the popup to automatically resize itself to prevent it from overflowing.",
							"value": bson.M{
								"type": "'horizontal' | 'vertical' | 'both'",
							},
						},
						bson.M{
							"name":        "sync",
							"description": "Syncs the popup's width or height to that of the anchor element.",
							"value": bson.M{
								"type": "'width' | 'height' | 'both'",
							},
						},
						bson.M{
							"name":        "autoSizeBoundary",
							"description": "The auto-size boundary describes clipping element(s) that overflow will be checked relative to when resizing. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
							"value": bson.M{
								"type": "Element | Elementbson.A{}",
							},
						},
						bson.M{
							"name":        "auto-size-padding",
							"description": "The amount of padding, in pixels, to exceed before the auto-size behavior will occur.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "hover-bridge",
							"description": "When a gap exists between the anchor and the popup element, this option will add a \"hover bridge\" that fills the\ngap using an invisible element. This makes listening for events such as `mouseenter` and `mouseleave` more sane\nbecause the pointer never technically leaves the element. The hover bridge will only be drawn when the popover is\nactive.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The popup's content.",
						},
						bson.M{
							"name":        "anchor",
							"description": "The element the popup will be anchored to. If the anchor lives outside of the popup, you can use the `anchor` attribute or property instead.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-reposition",
							"description": "Emitted when the popup is repositioned. This event can fire a lot, so avoid putting expensive operations in your listener or consider debouncing it.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "popup",
								"description": "A reference to the internal popup container. Useful for animating and styling the popup with JavaScript.",
								"type":        "HTMLElement",
							},
							bson.M{
								"name":        "anchor",
								"description": "The element the popup will be anchored to. If the anchor lives outside of the popup, you can provide the anchor\nelement `id`, a DOM element reference, or a `VirtualElement`. If the anchor lives inside the popup, use the\n`anchor` slot instead.",
								"type":        "Element | string | VirtualElement",
							},
							bson.M{
								"name":        "active",
								"description": "Activates the positioning logic and shows the popup. When this attribute is removed, the positioning logic is torn\ndown and the popup will be hidden.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "placement",
								"description": "The preferred placement of the popup. Note that the actual placement will vary as configured to keep the\npanel inside of the viewport.",
								"type":        "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
							},
							bson.M{
								"name":        "strategy",
								"description": "Determines how the popup is positioned. The `absolute` strategy works well in most cases, but if overflow is\nclipped, using a `fixed` position strategy can often workaround it.",
								"type":        "'absolute' | 'fixed'",
							},
							bson.M{
								"name":        "distance",
								"description": "The distance in pixels from which to offset the panel away from its anchor.",
								"type":        "number",
							},
							bson.M{
								"name":        "skidding",
								"description": "The distance in pixels from which to offset the panel along its anchor.",
								"type":        "number",
							},
							bson.M{
								"name":        "arrow",
								"description": "Attaches an arrow to the popup. The arrow's size and color can be customized using the `--arrow-size` and\n`--arrow-color` custom properties. For additional customizations, you can also target the arrow using\n`::part(arrow)` in your stylesheet.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "arrowPlacement",
								"description": "The placement of the arrow. The default is `anchor`, which will align the arrow as close to the center of the\nanchor as possible, considering available space and `arrow-padding`. A value of `start`, `end`, or `center` will\nalign the arrow to the start, end, or center of the popover instead.",
								"type":        "'start' | 'end' | 'center' | 'anchor'",
							},
							bson.M{
								"name":        "arrowPadding",
								"description": "The amount of padding between the arrow and the edges of the popup. If the popup has a border-radius, for example,\nthis will prevent it from overflowing the corners.",
								"type":        "number",
							},
							bson.M{
								"name":        "flip",
								"description": "When set, placement of the popup will flip to the opposite site to keep it in view. You can use\n`flipFallbackPlacements` to further configure how the fallback placement is determined.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "flipFallbackPlacements",
								"description": "If the preferred placement doesn't fit, popup will be tested in these fallback placements until one fits. Must be a\nstring of any number of placements separated by a space, e.g. \"top bottom left\". If no placement fits, the flip\nfallback strategy will be used instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "flipFallbackStrategy",
								"description": "When neither the preferred placement nor the fallback placements fit, this value will be used to determine whether\nthe popup should be positioned using the best available fit based on available space or as it was initially\npreferred.",
								"type":        "'best-fit' | 'initial'",
							},
							bson.M{
								"name":        "flipBoundary",
								"description": "The flip boundary describes clipping element(s) that overflow will be checked relative to when flipping. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
								"type":        "Element | Elementbson.A{}",
							},
							bson.M{
								"name":        "flipPadding",
								"description": "The amount of padding, in pixels, to exceed before the flip behavior will occur.",
								"type":        "number",
							},
							bson.M{
								"name":        "shift",
								"description": "Moves the popup along the axis to keep it in view when clipped.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "shiftBoundary",
								"description": "The shift boundary describes clipping element(s) that overflow will be checked relative to when shifting. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
								"type":        "Element | Elementbson.A{}",
							},
							bson.M{
								"name":        "shiftPadding",
								"description": "The amount of padding, in pixels, to exceed before the shift behavior will occur.",
								"type":        "number",
							},
							bson.M{
								"name":        "autoSize",
								"description": "When set, this will cause the popup to automatically resize itself to prevent it from overflowing.",
								"type":        "'horizontal' | 'vertical' | 'both'",
							},
							bson.M{
								"name":        "sync",
								"description": "Syncs the popup's width or height to that of the anchor element.",
								"type":        "'width' | 'height' | 'both'",
							},
							bson.M{
								"name":        "autoSizeBoundary",
								"description": "The auto-size boundary describes clipping element(s) that overflow will be checked relative to when resizing. By\ndefault, the boundary includes overflow ancestors that will cause the element to be clipped. If needed, you can\nchange the boundary by passing a reference to one or more elements to this property.",
								"type":        "Element | Elementbson.A{}",
							},
							bson.M{
								"name":        "autoSizePadding",
								"description": "The amount of padding, in pixels, to exceed before the auto-size behavior will occur.",
								"type":        "number",
							},
							bson.M{
								"name":        "hoverBridge",
								"description": "When a gap exists between the anchor and the popup element, this option will add a \"hover bridge\" that fills the\ngap using an invisible element. This makes listening for events such as `mouseenter` and `mouseleave` more sane\nbecause the pointer never technically leaves the element. The hover bridge will only be drawn when the popover is\nactive.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-reposition",
								"description": "Emitted when the popup is repositioned. This event can fire a lot, so avoid putting expensive operations in your listener or consider debouncing it.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-progress-bar",
					"description": "Progress bars are used to show the status of an ongoing operation.\n---\n\n\n### **Slots:**\n - _default_ - A label to show inside the progress indicator.\n\n### **CSS Properties:**\n - **--height** - The progress bar's height. _(default: undefined)_\n- **--track-color** - The color of the track. _(default: undefined)_\n- **--indicator-color** - The color of the indicator. _(default: undefined)_\n- **--label-color** - The color of the label. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **indicator** - The progress bar's indicator.\n- **label** - The progress bar's label.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The current progress as a percentage, 0 to 100.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "indeterminate",
							"description": "When true, percentage is ignored, the label is hidden, and the progress bar is drawn in an indeterminate state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "A custom label for assistive devices.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "A label to show inside the progress indicator.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "value",
								"description": "The current progress as a percentage, 0 to 100.",
								"type":        "number",
							},
							bson.M{
								"name":        "indeterminate",
								"description": "When true, percentage is ignored, the label is hidden, and the progress bar is drawn in an indeterminate state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "A custom label for assistive devices.",
								"type":        "string",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-progress-ring",
					"description": "Progress rings are used to show the progress of a determinate operation in a circular fashion.\n---\n\n\n### **Slots:**\n - _default_ - A label to show inside the ring.\n\n### **CSS Properties:**\n - **--size** - The diameter of the progress ring (cannot be a percentage). _(default: undefined)_\n- **--track-width** - The width of the track. _(default: undefined)_\n- **--track-color** - The color of the track. _(default: undefined)_\n- **--indicator-width** - The width of the indicator. Defaults to the track width. _(default: undefined)_\n- **--indicator-color** - The color of the indicator. _(default: undefined)_\n- **--indicator-transition-duration** - The duration of the indicator's transition when the value changes. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **label** - The progress ring label.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The current progress as a percentage, 0 to 100.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "label",
							"description": "A custom label for assistive devices.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "A label to show inside the ring.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "indicator",
								"type": "SVGCircleElement",
							},
							bson.M{
								"name": "indicatorOffset",
								"type": "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current progress as a percentage, 0 to 100.",
								"type":        "number",
							},
							bson.M{
								"name":        "label",
								"description": "A custom label for assistive devices.",
								"type":        "string",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-qr-code",
					"description": "Generates a bson.A{QR code}(https://www.qrcode.com/) and renders it using the bson.A{Canvas API}(https://developer.mozilla.org/en-US/docs/Web/API/Canvas_API).\n---\n\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The QR code's value.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The label for assistive devices to announce. If unspecified, the value will be used instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The size of the QR code, in pixels.",
							"value": bson.M{
								"type":    "number",
								"default": "128",
							},
						},
						bson.M{
							"name":        "fill",
							"description": "The fill color. This can be any valid CSS color, but not a CSS custom property.",
							"value": bson.M{
								"type":    "string",
								"default": "'black'",
							},
						},
						bson.M{
							"name":        "background",
							"description": "The background color. This can be any valid CSS color or `transparent`. It cannot be a CSS custom property.",
							"value": bson.M{
								"type":    "string",
								"default": "'white'",
							},
						},
						bson.M{
							"name":        "radius",
							"description": "The edge radius of each module. Must be between 0 and 0.5.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "error-correction",
							"description": "The level of error correction to use. bson.A{Learn more}(https://www.qrcode.com/en/about/error_correction.html)",
							"value": bson.M{
								"type":    "'L' | 'M' | 'Q' | 'H'",
								"default": "'H'",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "canvas",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "value",
								"description": "The QR code's value.",
								"type":        "string",
							},
							bson.M{
								"name":        "label",
								"description": "The label for assistive devices to announce. If unspecified, the value will be used instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The size of the QR code, in pixels.",
								"type":        "number",
							},
							bson.M{
								"name":        "fill",
								"description": "The fill color. This can be any valid CSS color, but not a CSS custom property.",
								"type":        "string",
							},
							bson.M{
								"name":        "background",
								"description": "The background color. This can be any valid CSS color or `transparent`. It cannot be a CSS custom property.",
								"type":        "string",
							},
							bson.M{
								"name":        "radius",
								"description": "The edge radius of each module. Must be between 0 and 0.5.",
								"type":        "number",
							},
							bson.M{
								"name":        "errorCorrection",
								"description": "The level of error correction to use. bson.A{Learn more}(https://www.qrcode.com/en/about/error_correction.html)",
								"type":        "'L' | 'M' | 'Q' | 'H'",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-radio",
					"description": "Radios allow the user to select a single option from a group.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the control loses focus.\n- **sl-focus** - Emitted when the control gains focus.\n\n### **Slots:**\n - _default_ - The radio's label.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **control** - The circular container that wraps the radio's checked state.\n- **control--checked** - The radio control when the radio is checked.\n- **checked-icon** - The checked icon, an `<sl-icon>` element.\n- **label** - The container that wraps the radio's label.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The radio's value. When selected, the radio group will receive this value.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The radio's size. When used inside a radio group, the size will be determined by the radio group's size so this\nattribute can typically be omitted.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the radio.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The radio's label.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "checked",
								"type": "boolean",
							},
							bson.M{
								"name":        "value",
								"description": "The radio's value. When selected, the radio group will receive this value.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The radio's size. When used inside a radio group, the size will be determined by the radio group's size so this\nattribute can typically be omitted.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the radio.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-menu-item",
					"description": "Menu items provide options for the user to pick from in a menu.\n---\n\n\n### **Methods:**\n - **getTextLabel()** - Returns a text label based on the contents of the menu item's default slot.\n\n### **Slots:**\n - _default_ - The menu item's label.\n- **prefix** - Used to prepend an icon or similar element to the menu item.\n- **suffix** - Used to append an icon or similar element to the menu item.\n- **submenu** - Used to denote a nested menu.\n\n### **CSS Properties:**\n - **--submenu-offset** - The distance submenus shift to overlap the parent menu. _(default: -2px)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **checked-icon** - The checked icon, which is only visible when the menu item is checked.\n- **prefix** - The prefix container.\n- **label** - The menu item label.\n- **suffix** - The suffix container.\n- **spinner** - The spinner that shows when the menu item is in the loading state.\n- **spinner__base** - The spinner's base part.\n- **submenu-icon** - The submenu icon, visible only when the menu item has a submenu (not yet implemented).",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "type",
							"description": "The type of menu item to render. To use `checked`, this value must be set to `checkbox`.",
							"value": bson.M{
								"type":    "'normal' | 'checkbox'",
								"default": "'normal'",
							},
						},
						bson.M{
							"name":        "checked",
							"description": "Draws the item in a checked state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "value",
							"description": "A unique value to store in the menu item. This can be used as a way to identify menu items when selected.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "loading",
							"description": "Draws the menu item in a loading state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Draws the menu item in a disabled state, preventing selection.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The menu item's label.",
						},
						bson.M{
							"name":        "prefix",
							"description": "Used to prepend an icon or similar element to the menu item.",
						},
						bson.M{
							"name":        "suffix",
							"description": "Used to append an icon or similar element to the menu item.",
						},
						bson.M{
							"name":        "submenu",
							"description": "Used to denote a nested menu.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "menuItem",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "type",
								"description": "The type of menu item to render. To use `checked`, this value must be set to `checkbox`.",
								"type":        "'normal' | 'checkbox'",
							},
							bson.M{
								"name":        "checked",
								"description": "Draws the item in a checked state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "value",
								"description": "A unique value to store in the menu item. This can be used as a way to identify menu items when selected.",
								"type":        "string",
							},
							bson.M{
								"name":        "loading",
								"description": "Draws the menu item in a loading state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Draws the menu item in a disabled state, preventing selection.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-radio-button",
					"description": "Radios buttons allow the user to select a single option from a group using a button-like control.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the button loses focus.\n- **sl-focus** - Emitted when the button gains focus.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the radio button.\n- **blur()** - Removes focus from the radio button.\n\n### **Slots:**\n - _default_ - The radio button's label.\n- **prefix** - A presentational prefix icon or similar element.\n- **suffix** - A presentational suffix icon or similar element.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **button** - The internal `<button>` element.\n- **button--checked** - The internal button element when the radio button is checked.\n- **prefix** - The container that wraps the prefix.\n- **label** - The container that wraps the radio button's label.\n- **suffix** - The container that wraps the suffix.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "value",
							"description": "The radio's value. When selected, the radio group will receive this value.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the radio button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The radio button's size. When used inside a radio group, the size will be determined by the radio group's size so\nthis attribute can typically be omitted.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style radio button with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The radio button's label.",
						},
						bson.M{
							"name":        "prefix",
							"description": "A presentational prefix icon or similar element.",
						},
						bson.M{
							"name":        "suffix",
							"description": "A presentational suffix icon or similar element.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the button loses focus.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the button gains focus.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "hiddenInput",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name":        "value",
								"description": "The radio's value. When selected, the radio group will receive this value.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the radio button.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "size",
								"description": "The radio button's size. When used inside a radio group, the size will be determined by the radio group's size so\nthis attribute can typically be omitted.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style radio button with rounded edges.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the button loses focus.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the button gains focus.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-radio-group",
					"description": "Radio groups are used to group multiple bson.A{radios}(/components/radio) or bson.A{radio buttons}(/components/radio-button) so they function as a single form control.\n---\n\n\n### **Events:**\n - **sl-change** - Emitted when the radio group's selected value changes.\n- **sl-input** - Emitted when the radio group receives user input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity(): _boolean_** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - _default_ - The default slot where `<sl-radio>` or `<sl-radio-button>` elements are placed.\n- **label** - The radio group's label. Required for proper accessibility. Alternatively, you can use the `label` attribute.\n- **help-text** - Text that describes how to use the radio group. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Parts:**\n - **form-control** - The form control that wraps the label, input, and help text.\n- **form-control-label** - The label's wrapper.\n- **form-control-input** - The input's wrapper.\n- **form-control-help-text** - The help text's wrapper.\n- **button-group** - The button group that wraps radio buttons.\n- **button-group__base** - The button group's `base` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "label",
							"description": "The radio group's label. Required for proper accessibility. If you need to display HTML, use the `label` slot\ninstead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The radio groups's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the radio group, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "'option'",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the radio group, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The radio group's size. This size will be applied to all child radios and radio buttons.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Ensures a child radio is checked before allowing the containing form to submit.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The default slot where `<sl-radio>` or `<sl-radio-button>` elements are placed.",
						},
						bson.M{
							"name":        "label",
							"description": "The radio group's label. Required for proper accessibility. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the radio group. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the radio group's selected value changes.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the radio group receives user input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "validationInput",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "defaultValue",
								"type": "string",
							},
							bson.M{
								"name":        "label",
								"description": "The radio group's label. Required for proper accessibility. If you need to display HTML, use the `label` slot\ninstead.",
								"type":        "string",
							},
							bson.M{
								"name":        "helpText",
								"description": "The radio groups's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the radio group, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the radio group, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The radio group's size. This size will be applied to all child radios and radio buttons.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Ensures a child radio is checked before allowing the containing form to submit.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the radio group's selected value changes.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the radio group receives user input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-range",
					"description": "Ranges allow the user to select a single value within a given range using a slider.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the control loses focus.\n- **sl-change** - Emitted when an alteration to the control's value is committed by the user.\n- **sl-focus** - Emitted when the control gains focus.\n- **sl-input** - Emitted when the control receives input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the range.\n- **blur()** - Removes focus from the range.\n- **stepUp()** - Increments the value of the range by the value of the step attribute.\n- **stepDown()** - Decrements the value of the range by the value of the step attribute.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - **label** - The range's label. Alternatively, you can use the `label` attribute.\n- **help-text** - Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Properties:**\n - **--thumb-size** - The size of the thumb. _(default: undefined)_\n- **--tooltip-offset** - The vertical distance the tooltip is offset from the track. _(default: undefined)_\n- **--track-color-active** - The color of the portion of the track that represents the current value. _(default: undefined)_\n- **--track-color-inactive** - The of the portion of the track that represents the remaining value. _(default: undefined)_\n- **--track-height** - The height of the track. _(default: undefined)_\n- **--track-active-offset** - The point of origin of the active track. _(default: undefined)_\n\n### **CSS Parts:**\n - **form-control** - The form control that wraps the label, input, and help text.\n- **form-control-label** - The label's wrapper.\n- **form-control-input** - The range's wrapper.\n- **form-control-help-text** - The help text's wrapper.\n- **base** - The component's base wrapper.\n- **input** - The internal `<input>` element.\n- **tooltip** - The range's tooltip.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the range, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the range, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The range's label. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The range's help text. If you need to display HTML, use the help-text slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the range.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "min",
							"description": "The minimum acceptable value of the range.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "max",
							"description": "The maximum acceptable value of the range.",
							"value": bson.M{
								"type":    "number",
								"default": "100",
							},
						},
						bson.M{
							"name":        "step",
							"description": "The interval at which the range will increase and decrease.",
							"value": bson.M{
								"type":    "number",
								"default": "1",
							},
						},
						bson.M{
							"name":        "tooltip",
							"description": "The preferred placement of the range's tooltip.",
							"value": bson.M{
								"type":    "'top' | 'bottom' | 'none'",
								"default": "'top'",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "label",
							"description": "The range's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when an alteration to the control's value is committed by the user.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the control receives input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "output",
								"type": "HTMLOutputElement | null",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the range, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the range, submitted as a name/value pair with form data.",
								"type":        "number",
							},
							bson.M{
								"name":        "label",
								"description": "The range's label. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "helpText",
								"description": "The range's help text. If you need to display HTML, use the help-text slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the range.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "min",
								"description": "The minimum acceptable value of the range.",
								"type":        "number",
							},
							bson.M{
								"name":        "max",
								"description": "The maximum acceptable value of the range.",
								"type":        "number",
							},
							bson.M{
								"name":        "step",
								"description": "The interval at which the range will increase and decrease.",
								"type":        "number",
							},
							bson.M{
								"name":        "tooltip",
								"description": "The preferred placement of the range's tooltip.",
								"type":        "'top' | 'bottom' | 'none'",
							},
							bson.M{
								"name":        "tooltipFormatter",
								"description": "A function used to format the tooltip's value. The range's value is passed as the first and only argument. The\nfunction should return a string to display in the tooltip.",
								"type":        "(value: number) => string",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "defaultValue",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "number",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when an alteration to the control's value is committed by the user.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the control receives input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-relative-time",
					"description": "Outputs a localized time phrase relative to the current date and time.\n---\n",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "date",
							"description": "The date from which to calculate time from. If not set, the current date and time will be used. When passing a\nstring, it's strongly recommended to use the ISO 8601 format to ensure timezones are handled correctly. To convert\na date to this format in JavaScript, use bson.A{`date.toISOString()`}(https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString).",
							"value": bson.M{
								"type":    "Date | string",
								"default": "new Date()",
							},
						},
						bson.M{
							"name":        "format",
							"description": "The formatting style to use.",
							"value": bson.M{
								"type":    "'long' | 'short' | 'narrow'",
								"default": "'long'",
							},
						},
						bson.M{
							"name":        "numeric",
							"description": "When `auto`, values such as \"yesterday\" and \"tomorrow\" will be shown when possible. When `always`, values such as\n\"1 day ago\" and \"in 1 day\" will be shown.",
							"value": bson.M{
								"type":    "'always' | 'auto'",
								"default": "'auto'",
							},
						},
						bson.M{
							"name":        "sync",
							"description": "Keep the displayed value up to date as time passes.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "date",
								"description": "The date from which to calculate time from. If not set, the current date and time will be used. When passing a\nstring, it's strongly recommended to use the ISO 8601 format to ensure timezones are handled correctly. To convert\na date to this format in JavaScript, use bson.A{`date.toISOString()`}(https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString).",
								"type":        "Date | string",
							},
							bson.M{
								"name":        "format",
								"description": "The formatting style to use.",
								"type":        "'long' | 'short' | 'narrow'",
							},
							bson.M{
								"name":        "numeric",
								"description": "When `auto`, values such as \"yesterday\" and \"tomorrow\" will be shown when possible. When `always`, values such as\n\"1 day ago\" and \"in 1 day\" will be shown.",
								"type":        "'always' | 'auto'",
							},
							bson.M{
								"name":        "sync",
								"description": "Keep the displayed value up to date as time passes.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-resize-observer",
					"description": "The Resize Observer component offers a thin, declarative interface to the bson.A{`ResizeObserver API`}(https://developer.mozilla.org/en-US/docs/Web/API/ResizeObserver).\n---\n\n\n### **Events:**\n - **sl-resize** - Emitted when the element is resized.\n\n### **Slots:**\n - _default_ - One or more elements to watch for resizing.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "disabled",
							"description": "Disables the observer.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "One or more elements to watch for resizing.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-resize",
							"type":        "bson.M{ entries: ResizeObserverEntrybson.A{} }",
							"description": "Emitted when the element is resized.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "disabled",
								"description": "Disables the observer.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-resize",
								"type":        "bson.M{ entries: ResizeObserverEntrybson.A{} }",
								"description": "Emitted when the element is resized.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-select",
					"description": "Selects allow you to choose items from a menu of predefined options.\n---\n\n\n### **Events:**\n - **sl-change** - Emitted when the control's value changes.\n- **sl-clear** - Emitted when the control's value is cleared.\n- **sl-input** - Emitted when the control receives input.\n- **sl-focus** - Emitted when the control gains focus.\n- **sl-blur** - Emitted when the control loses focus.\n- **sl-show** - Emitted when the select's menu opens.\n- **sl-after-show** - Emitted after the select's menu opens and all animations are complete.\n- **sl-hide** - Emitted when the select's menu closes.\n- **sl-after-hide** - Emitted after the select's menu closes and all animations are complete.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **show()** - Shows the listbox.\n- **hide()** - Hides the listbox.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n- **focus(options: _FocusOptions_)** - Sets focus on the control.\n- **blur()** - Removes focus from the control.\n\n### **Slots:**\n - _default_ - The listbox options. Must be `<sl-option>` elements. You can use `<sl-divider>` to group items visually.\n- **label** - The input's label. Alternatively, you can use the `label` attribute.\n- **prefix** - Used to prepend a presentational icon or similar element to the combobox.\n- **clear-icon** - An icon to use in lieu of the default clear icon.\n- **expand-icon** - The icon to show when the control is expanded and collapsed. Rotates on open and close.\n- **help-text** - Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Parts:**\n - **form-control** - The form control that wraps the label, input, and help text.\n- **form-control-label** - The label's wrapper.\n- **form-control-input** - The select's wrapper.\n- **form-control-help-text** - The help text's wrapper.\n- **combobox** - The container the wraps the prefix, combobox, clear icon, and expand button.\n- **prefix** - The container that wraps the prefix slot.\n- **display-input** - The element that displays the selected option's label, an `<input>` element.\n- **listbox** - The listbox container where options are slotted.\n- **tags** - The container that houses option tags when `multiselect` is used.\n- **tag** - The individual tags that represent each multiselect option.\n- **tag__base** - The tag's base part.\n- **tag__content** - The tag's content part.\n- **tag__remove-button** - The tag's remove button.\n- **tag__remove-button__base** - The tag's remove button base part.\n- **clear-button** - The clear button.\n- **expand-icon** - The container that wraps the expand icon.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "name",
							"description": "The name of the select, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the select, submitted as a name/value pair with form data. When `multiple` is enabled, the\nvalue attribute will be a space-delimited list of values based on the options selected, and the value property will\nbe an array. **For this reason, values must not contain spaces.**",
							"value": bson.M{
								"type":    "string | stringbson.A{}",
								"default": "''",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The select's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "placeholder",
							"description": "Placeholder text to show as a hint when the select is empty.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "multiple",
							"description": "Allows more than one option to be selected.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "max-options-visible",
							"description": "The maximum number of selected options to show when `multiple` is true. After the maximum, \"+n\" will be shown to\nindicate the number of additional items that are selected. Set to 0 to remove the limit.",
							"value": bson.M{
								"type":    "number",
								"default": "3",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the select control.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "clearable",
							"description": "Adds a clear button when the select is not empty.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the select is open. You can toggle this attribute to show and hide the menu, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the select's open state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "hoist",
							"description": "Enable this option to prevent the listbox from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "filled",
							"description": "Draws a filled select.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style select with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The select's label. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "placement",
							"description": "The preferred placement of the select's menu. Note that the actual placement may vary as needed to keep the listbox\ninside of the viewport.",
							"value": bson.M{
								"type":    "'top' | 'bottom'",
								"default": "'bottom'",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The select's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "The select's required attribute.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "getTag",
							"description": "A function that customizes the tags to be rendered when multiple=true. The first argument is the option, the second\nis the current tag's index.  The function should return either a Lit TemplateResult or a string containing trusted HTML of the symbol to render at\nthe specified value.",
							"value": bson.M{
								"type": "(option: SlOption, index: number) => TemplateResult | string | HTMLElement",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The listbox options. Must be `<sl-option>` elements. You can use `<sl-divider>` to group items visually.",
						},
						bson.M{
							"name":        "label",
							"description": "The input's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "prefix",
							"description": "Used to prepend a presentational icon or similar element to the combobox.",
						},
						bson.M{
							"name":        "clear-icon",
							"description": "An icon to use in lieu of the default clear icon.",
						},
						bson.M{
							"name":        "expand-icon",
							"description": "The icon to show when the control is expanded and collapsed. Rotates on open and close.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the control's value changes.",
						},
						bson.M{
							"name":        "sl-clear",
							"description": "Emitted when the control's value is cleared.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the control receives input.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the select's menu opens.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the select's menu opens and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the select's menu closes.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the select's menu closes and all animations are complete.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "popup",
								"type": "SlPopup",
							},
							bson.M{
								"name": "combobox",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "displayInput",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "valueInput",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "listbox",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "displayLabel",
								"type": "string",
							},
							bson.M{
								"name": "currentOption",
								"type": "SlOption",
							},
							bson.M{
								"name": "selectedOptions",
								"type": "SlOptionbson.A{}",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the select, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the select, submitted as a name/value pair with form data. When `multiple` is enabled, the\nvalue attribute will be a space-delimited list of values based on the options selected, and the value property will\nbe an array. **For this reason, values must not contain spaces.**",
								"type":        "string | stringbson.A{}",
							},
							bson.M{
								"name":        "defaultValue",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "string | stringbson.A{}",
							},
							bson.M{
								"name":        "size",
								"description": "The select's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "placeholder",
								"description": "Placeholder text to show as a hint when the select is empty.",
								"type":        "string",
							},
							bson.M{
								"name":        "multiple",
								"description": "Allows more than one option to be selected.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "maxOptionsVisible",
								"description": "The maximum number of selected options to show when `multiple` is true. After the maximum, \"+n\" will be shown to\nindicate the number of additional items that are selected. Set to 0 to remove the limit.",
								"type":        "number",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the select control.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "clearable",
								"description": "Adds a clear button when the select is not empty.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the select is open. You can toggle this attribute to show and hide the menu, or you can\nuse the `show()` and `hide()` methods and this attribute will reflect the select's open state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "hoist",
								"description": "Enable this option to prevent the listbox from being clipped when the component is placed inside a container with\n`overflow: auto|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all, scenarios.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "filled",
								"description": "Draws a filled select.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style select with rounded edges.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "The select's label. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "placement",
								"description": "The preferred placement of the select's menu. Note that the actual placement may vary as needed to keep the listbox\ninside of the viewport.",
								"type":        "'top' | 'bottom'",
							},
							bson.M{
								"name":        "helpText",
								"description": "The select's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "The select's required attribute.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "getTag",
								"description": "A function that customizes the tags to be rendered when multiple=true. The first argument is the option, the second\nis the current tag's index.  The function should return either a Lit TemplateResult or a string containing trusted HTML of the symbol to render at\nthe specified value.",
								"type":        "(option: SlOption, index: number) => TemplateResult | string | HTMLElement",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the control's value changes.",
							},
							bson.M{
								"name":        "sl-clear",
								"description": "Emitted when the control's value is cleared.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the control receives input.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the select's menu opens.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the select's menu opens and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the select's menu closes.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the select's menu closes and all animations are complete.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-skeleton",
					"description": "Skeletons are used to provide a visual representation of where content will eventually be drawn.\n---\n\n\n### **CSS Properties:**\n - **--border-radius** - The skeleton's border radius. _(default: undefined)_\n- **--color** - The color of the skeleton. _(default: undefined)_\n- **--sheen-color** - The sheen color when the skeleton is in its loading state. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **indicator** - The skeleton's indicator which is responsible for its color and animation.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "effect",
							"description": "Determines which effect the skeleton will use.",
							"value": bson.M{
								"type":    "'pulse' | 'sheen' | 'none'",
								"default": "'none'",
							},
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "effect",
								"description": "Determines which effect the skeleton will use.",
								"type":        "'pulse' | 'sheen' | 'none'",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-rating",
					"description": "Ratings give users a way to quickly view and provide feedback.\n---\n\n\n### **Events:**\n - **sl-change** - Emitted when the rating's value changes.\n- **sl-hover** - Emitted when the user hovers over a value. The `phase` property indicates when hovering starts, moves to a new value, or ends. The `value` property tells what the rating's value would be if the user were to commit to the hovered value.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the rating.\n- **blur()** - Removes focus from the rating.\n\n### **CSS Properties:**\n - **--symbol-color** - The inactive color for symbols. _(default: undefined)_\n- **--symbol-color-active** - The active color for symbols. _(default: undefined)_\n- **--symbol-size** - The size of symbols. _(default: undefined)_\n- **--symbol-spacing** - The spacing to use around symbols. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "label",
							"description": "A label that describes the rating to assistive devices.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current rating.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "max",
							"description": "The highest rating to show.",
							"value": bson.M{
								"type":    "number",
								"default": "5",
							},
						},
						bson.M{
							"name":        "precision",
							"description": "The precision at which the rating will increase and decrease. For example, to allow half-star ratings, set this\nattribute to `0.5`.",
							"value": bson.M{
								"type":    "number",
								"default": "1",
							},
						},
						bson.M{
							"name":        "readonly",
							"description": "Makes the rating readonly.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the rating.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "getSymbol",
							"description": "A function that customizes the symbol to be rendered. The first and only argument is the rating's current value.\nThe function should return a string containing trusted HTML of the symbol to render at the specified value. Works\nwell with `<sl-icon>` elements.",
							"value": bson.M{
								"type": "(value: number) => string",
							},
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the rating's value changes.",
						},
						bson.M{
							"name":        "sl-hover",
							"type":        "bson.M{ phase: 'start' | 'move' | 'end', value: number }",
							"description": "Emitted when the user hovers over a value. The `phase` property indicates when hovering starts, moves to a new value, or ends. The `value` property tells what the rating's value would be if the user were to commit to the hovered value.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "rating",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "label",
								"description": "A label that describes the rating to assistive devices.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current rating.",
								"type":        "number",
							},
							bson.M{
								"name":        "max",
								"description": "The highest rating to show.",
								"type":        "number",
							},
							bson.M{
								"name":        "precision",
								"description": "The precision at which the rating will increase and decrease. For example, to allow half-star ratings, set this\nattribute to `0.5`.",
								"type":        "number",
							},
							bson.M{
								"name":        "readonly",
								"description": "Makes the rating readonly.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the rating.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "getSymbol",
								"description": "A function that customizes the symbol to be rendered. The first and only argument is the rating's current value.\nThe function should return a string containing trusted HTML of the symbol to render at the specified value. Works\nwell with `<sl-icon>` elements.",
								"type":        "(value: number) => string",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the rating's value changes.",
							},
							bson.M{
								"name":        "sl-hover",
								"type":        "bson.M{ phase: 'start' | 'move' | 'end', value: number }",
								"description": "Emitted when the user hovers over a value. The `phase` property indicates when hovering starts, moves to a new value, or ends. The `value` property tells what the rating's value would be if the user were to commit to the hovered value.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-spinner",
					"description": "Spinners are used to show the progress of an indeterminate operation.\n---\n\n\n### **CSS Properties:**\n - **--track-width** - The width of the track. _(default: undefined)_\n- **--track-color** - The color of the track. _(default: undefined)_\n- **--indicator-color** - The color of the spinner's indicator. _(default: undefined)_\n- **--speed** - The time it takes for the spinner to complete one animation cycle. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes":  bson.A{},
					"events":      bson.A{},
					"js": bson.M{
						"properties": bson.A{},
						"events":     bson.A{},
					},
				},
				bson.M{
					"name":        "sl-split-panel",
					"description": "Split panels display two adjacent panels, allowing the user to reposition them.\n---\n\n\n### **Events:**\n - **sl-reposition** - Emitted when the divider's position changes.\n\n### **Slots:**\n - **start** - Content to place in the start panel.\n- **end** - Content to place in the end panel.\n- **divider** - The divider. Useful for slotting in a custom icon that renders as a handle.\n\n### **CSS Properties:**\n - **--divider-width** - The width of the visible divider. _(default: 4px)_\n- **--divider-hit-area** - The invisible region around the divider where dragging can occur. This is usually wider than the divider to facilitate easier dragging. _(default: 12px)_\n- **--min** - The minimum allowed size of the primary panel. _(default: 0)_\n- **--max** - The maximum allowed size of the primary panel. _(default: 100%)_\n\n### **CSS Parts:**\n - **start** - The start panel.\n- **end** - The end panel.\n- **panel** - Targets both the start and end panels.\n- **divider** - The divider that separates the start and end panels.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "position",
							"description": "The current position of the divider from the primary panel's edge as a percentage 0-100. Defaults to 50% of the\ncontainer's initial size.",
							"value": bson.M{
								"type":    "number",
								"default": "50",
							},
						},
						bson.M{
							"name":        "position-in-pixels",
							"description": "The current position of the divider from the primary panel's edge in pixels.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "vertical",
							"description": "Draws the split panel in a vertical orientation with the start and end panels stacked.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables resizing. Note that the position may still change as a result of resizing the host element.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "primary",
							"description": "If no primary panel is designated, both panels will resize proportionally when the host element is resized. If a\nprimary panel is designated, it will maintain its size and the other panel will grow or shrink as needed when the\nhost element is resized.",
							"value": bson.M{
								"type": "'start' | 'end' | undefined",
							},
						},
						bson.M{
							"name":        "snap",
							"description": "One or more space-separated values at which the divider should snap. Values can be in pixels or percentages, e.g.\n`\"100px 50%\"`.",
							"value": bson.M{
								"type": "string | undefined",
							},
						},
						bson.M{
							"name":        "snap-threshold",
							"description": "How close the divider must be to a snap point until snapping occurs.",
							"value": bson.M{
								"type":    "number",
								"default": "12",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "start",
							"description": "Content to place in the start panel.",
						},
						bson.M{
							"name":        "end",
							"description": "Content to place in the end panel.",
						},
						bson.M{
							"name":        "divider",
							"description": "The divider. Useful for slotting in a custom icon that renders as a handle.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-reposition",
							"description": "Emitted when the divider's position changes.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "divider",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "position",
								"description": "The current position of the divider from the primary panel's edge as a percentage 0-100. Defaults to 50% of the\ncontainer's initial size.",
								"type":        "number",
							},
							bson.M{
								"name":        "positionInPixels",
								"description": "The current position of the divider from the primary panel's edge in pixels.",
								"type":        "number",
							},
							bson.M{
								"name":        "vertical",
								"description": "Draws the split panel in a vertical orientation with the start and end panels stacked.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables resizing. Note that the position may still change as a result of resizing the host element.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "primary",
								"description": "If no primary panel is designated, both panels will resize proportionally when the host element is resized. If a\nprimary panel is designated, it will maintain its size and the other panel will grow or shrink as needed when the\nhost element is resized.",
								"type":        "'start' | 'end' | undefined",
							},
							bson.M{
								"name":        "snap",
								"description": "One or more space-separated values at which the divider should snap. Values can be in pixels or percentages, e.g.\n`\"100px 50%\"`.",
								"type":        "string | undefined",
							},
							bson.M{
								"name":        "snapThreshold",
								"description": "How close the divider must be to a snap point until snapping occurs.",
								"type":        "number",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-reposition",
								"description": "Emitted when the divider's position changes.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-switch",
					"description": "Switches allow the user to toggle an option on or off.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the control loses focus.\n- **sl-change** - Emitted when the control's checked state changes.\n- **sl-input** - Emitted when the control receives input.\n- **sl-focus** - Emitted when the control gains focus.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **click()** - Simulates a click on the switch.\n- **focus(options: _FocusOptions_)** - Sets focus on the switch.\n- **blur()** - Removes focus from the switch.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - _default_ - The switch's label.\n- **help-text** - Text that describes how to use the switch. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Properties:**\n - **--width** - The width of the switch. _(default: undefined)_\n- **--height** - The height of the switch. _(default: undefined)_\n- **--thumb-size** - The size of the thumb. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **control** - The control that houses the switch's thumb.\n- **thumb** - The switch's thumb.\n- **label** - The switch's label.\n- **form-control-help-text** - The help text's wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the switch, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the switch, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The switch's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the switch.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "checked",
							"description": "Draws the switch in a checked state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Makes the switch a required field.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The switch's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The switch's label.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the switch. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when the control's checked state changes.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the control receives input.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLInputElement",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the switch, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the switch, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The switch's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the switch.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "checked",
								"description": "Draws the switch in a checked state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "defaultChecked",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Makes the switch a required field.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "helpText",
								"description": "The switch's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when the control's checked state changes.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the control receives input.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-tab-group",
					"description": "Tab groups organize content into a container that shows one section at a time.\n---\n\n\n### **Events:**\n - **sl-tab-show** - Emitted when a tab is shown.\n- **sl-tab-hide** - Emitted when a tab is hidden.\n\n### **Methods:**\n - **show(panel: _string_)** - Shows the specified tab panel.\n\n### **Slots:**\n - _default_ - Used for grouping tab panels in the tab group. Must be `<sl-tab-panel>` elements.\n- **nav** - Used for grouping tabs in the tab group. Must be `<sl-tab>` elements.\n\n### **CSS Properties:**\n - **--indicator-color** - The color of the active tab indicator. _(default: undefined)_\n- **--track-color** - The color of the indicator's track (the line that separates tabs from panels). _(default: undefined)_\n- **--track-width** - The width of the indicator's track (the line that separates tabs from panels). _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **nav** - The tab group's navigation container where tabs are slotted in.\n- **tabs** - The container that wraps the tabs.\n- **active-tab-indicator** - The line that highlights the currently selected tab.\n- **body** - The tab group's body where tab panels are slotted in.\n- **scroll-button** - The previous/next scroll buttons that show when tabs are scrollable, an `<sl-icon-button>`.\n- **scroll-button--start** - The starting scroll button.\n- **scroll-button--end** - The ending scroll button.\n- **scroll-button__base** - The scroll button's exported `base` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "placement",
							"description": "The placement of the tabs.",
							"value": bson.M{
								"type":    "'top' | 'bottom' | 'start' | 'end'",
								"default": "'top'",
							},
						},
						bson.M{
							"name":        "activation",
							"description": "When set to auto, navigating tabs with the arrow keys will instantly show the corresponding tab panel. When set to\nmanual, the tab will receive focus but will not show until the user presses spacebar or enter.",
							"value": bson.M{
								"type":    "'auto' | 'manual'",
								"default": "'auto'",
							},
						},
						bson.M{
							"name":        "no-scroll-controls",
							"description": "Disables the scroll arrows that appear when tabs overflow.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "Used for grouping tab panels in the tab group. Must be `<sl-tab-panel>` elements.",
						},
						bson.M{
							"name":        "nav",
							"description": "Used for grouping tabs in the tab group. Must be `<sl-tab>` elements.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-tab-show",
							"type":        "bson.M{ name: String }",
							"description": "Emitted when a tab is shown.",
						},
						bson.M{
							"name":        "sl-tab-hide",
							"type":        "bson.M{ name: String }",
							"description": "Emitted when a tab is hidden.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "tabGroup",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "body",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "nav",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "indicator",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "placement",
								"description": "The placement of the tabs.",
								"type":        "'top' | 'bottom' | 'start' | 'end'",
							},
							bson.M{
								"name":        "activation",
								"description": "When set to auto, navigating tabs with the arrow keys will instantly show the corresponding tab panel. When set to\nmanual, the tab will receive focus but will not show until the user presses spacebar or enter.",
								"type":        "'auto' | 'manual'",
							},
							bson.M{
								"name":        "noScrollControls",
								"description": "Disables the scroll arrows that appear when tabs overflow.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-tab-show",
								"type":        "bson.M{ name: String }",
								"description": "Emitted when a tab is shown.",
							},
							bson.M{
								"name":        "sl-tab-hide",
								"type":        "bson.M{ name: String }",
								"description": "Emitted when a tab is hidden.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-tab-panel",
					"description": "Tab panels are used inside bson.A{tab groups}(/components/tab-group) to display tabbed content.\n---\n\n\n### **Slots:**\n - _default_ - The tab panel's content.\n\n### **CSS Properties:**\n - **--padding** - The tab panel's padding. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "name",
							"description": "The tab panel's name.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "active",
							"description": "When true, the tab panel will be shown.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The tab panel's content.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "name",
								"description": "The tab panel's name.",
								"type":        "string",
							},
							bson.M{
								"name":        "active",
								"description": "When true, the tab panel will be shown.",
								"type":        "boolean",
							},
						},
						"events": bson.A{},
					},
				},
				bson.M{
					"name":        "sl-tag",
					"description": "Tags are used as labels to organize things or to indicate a selection.\n---\n\n\n### **Events:**\n - **sl-remove** - Emitted when the remove button is activated.\n\n### **Slots:**\n - _default_ - The tag's content.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **content** - The tag's content.\n- **remove-button** - The tag's remove button, an `<sl-icon-button>`.\n- **remove-button__base** - The remove button's exported `base` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "variant",
							"description": "The tag's theme variant.",
							"value": bson.M{
								"type":    "'primary' | 'success' | 'neutral' | 'warning' | 'danger' | 'text'",
								"default": "'neutral'",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The tag's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "pill",
							"description": "Draws a pill-style tag with rounded edges.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "removable",
							"description": "Makes the tag removable and shows a remove button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The tag's content.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-remove",
							"description": "Emitted when the remove button is activated.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name":        "variant",
								"description": "The tag's theme variant.",
								"type":        "'primary' | 'success' | 'neutral' | 'warning' | 'danger' | 'text'",
							},
							bson.M{
								"name":        "size",
								"description": "The tag's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "pill",
								"description": "Draws a pill-style tag with rounded edges.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "removable",
								"description": "Makes the tag removable and shows a remove button.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-remove",
								"description": "Emitted when the remove button is activated.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-textarea",
					"description": "Textareas collect data from the user and allow multiple lines of text.\n---\n\n\n### **Events:**\n - **sl-blur** - Emitted when the control loses focus.\n- **sl-change** - Emitted when an alteration to the control's value is committed by the user.\n- **sl-focus** - Emitted when the control gains focus.\n- **sl-input** - Emitted when the control receives input.\n- **sl-invalid** - Emitted when the form control has been checked for validity and its constraints aren't satisfied.\n\n### **Methods:**\n - **focus(options: _FocusOptions_)** - Sets focus on the textarea.\n- **blur()** - Removes focus from the textarea.\n- **select()** - Selects all the text in the textarea.\n- **scrollPosition(position: _bson.M{ top?: number; left?: number }_): _bson.M{ top: number; left: number } | undefined_** - Gets or sets the textarea's scroll position.\n- **setSelectionRange(selectionStart: _number_, selectionEnd: _number_, selectionDirection: _'forward' | 'backward' | 'none'_)** - Sets the start and end positions of the text selection (0-based).\n- **setRangeText(replacement: _string_, start: _number_, end: _number_, selectMode: _'select' | 'start' | 'end' | 'preserve'_)** - Replaces a range of text with a new string.\n- **checkValidity()** - Checks for validity but does not show a validation message. Returns `true` when valid and `false` when invalid.\n- **getForm(): _HTMLFormElement | null_** - Gets the associated form, if one exists.\n- **reportValidity()** - Checks for validity and shows the browser's validation message if the control is invalid.\n- **setCustomValidity(message: _string_)** - Sets a custom validation message. Pass an empty string to restore validity.\n\n### **Slots:**\n - **label** - The textarea's label. Alternatively, you can use the `label` attribute.\n- **help-text** - Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.\n\n### **CSS Parts:**\n - **form-control** - The form control that wraps the label, input, and help text.\n- **form-control-label** - The label's wrapper.\n- **form-control-input** - The input's wrapper.\n- **form-control-help-text** - The help text's wrapper.\n- **base** - The component's base wrapper.\n- **textarea** - The internal `<textarea>` control.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name": "title",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "name",
							"description": "The name of the textarea, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "value",
							"description": "The current value of the textarea, submitted as a name/value pair with form data.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "size",
							"description": "The textarea's size.",
							"value": bson.M{
								"type":    "'small' | 'medium' | 'large'",
								"default": "'medium'",
							},
						},
						bson.M{
							"name":        "filled",
							"description": "Draws a filled textarea.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "label",
							"description": "The textarea's label. If you need to display HTML, use the `label` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "help-text",
							"description": "The textarea's help text. If you need to display HTML, use the `help-text` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "placeholder",
							"description": "Placeholder text to show as a hint when the input is empty.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "rows",
							"description": "The number of rows to display by default.",
							"value": bson.M{
								"type":    "number",
								"default": "4",
							},
						},
						bson.M{
							"name":        "resize",
							"description": "Controls how the textarea can be resized.",
							"value": bson.M{
								"type":    "'none' | 'vertical' | 'auto'",
								"default": "'vertical'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the textarea.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "readonly",
							"description": "Makes the textarea readonly.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "form",
							"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "required",
							"description": "Makes the textarea a required field.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "minlength",
							"description": "The minimum length of input that will be considered valid.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "maxlength",
							"description": "The maximum length of input that will be considered valid.",
							"value": bson.M{
								"type": "number",
							},
						},
						bson.M{
							"name":        "autocapitalize",
							"description": "Controls whether and how text input is automatically capitalized as it is entered by the user.",
							"value": bson.M{
								"type": "'off' | 'none' | 'on' | 'sentences' | 'words' | 'characters'",
							},
						},
						bson.M{
							"name":        "autocorrect",
							"description": "Indicates whether the browser's autocorrect feature is on or off.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "autocomplete",
							"description": "Specifies what permission the browser has to provide assistance in filling out form field values. Refer to\nbson.A{this page on MDN}(https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) for available values.",
							"value": bson.M{
								"type": "string",
							},
						},
						bson.M{
							"name":        "autofocus",
							"description": "Indicates that the input should receive focus on page load.",
							"value": bson.M{
								"type": "boolean",
							},
						},
						bson.M{
							"name":        "enterkeyhint",
							"description": "Used to customize the label or icon of the Enter key on virtual keyboards.",
							"value": bson.M{
								"type": "'enter' | 'done' | 'go' | 'next' | 'previous' | 'search' | 'send'",
							},
						},
						bson.M{
							"name":        "spellcheck",
							"description": "Enables spell checking on the textarea.",
							"value": bson.M{
								"type":    "boolean",
								"default": "true",
							},
						},
						bson.M{
							"name":        "inputmode",
							"description": "Tells the browser what type of data will be entered by the user, allowing it to display the appropriate virtual\nkeyboard on supportive devices.",
							"value": bson.M{
								"type": "'none' | 'text' | 'decimal' | 'numeric' | 'tel' | 'search' | 'email' | 'url'",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "label",
							"description": "The textarea's label. Alternatively, you can use the `label` attribute.",
						},
						bson.M{
							"name":        "help-text",
							"description": "Text that describes how to use the input. Alternatively, you can use the `help-text` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-blur",
							"description": "Emitted when the control loses focus.",
						},
						bson.M{
							"name":        "sl-change",
							"description": "Emitted when an alteration to the control's value is committed by the user.",
						},
						bson.M{
							"name":        "sl-focus",
							"description": "Emitted when the control gains focus.",
						},
						bson.M{
							"name":        "sl-input",
							"description": "Emitted when the control receives input.",
						},
						bson.M{
							"name":        "sl-invalid",
							"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "input",
								"type": "HTMLTextAreaElement",
							},
							bson.M{
								"name": "title",
								"type": "string",
							},
							bson.M{
								"name":        "name",
								"description": "The name of the textarea, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "value",
								"description": "The current value of the textarea, submitted as a name/value pair with form data.",
								"type":        "string",
							},
							bson.M{
								"name":        "size",
								"description": "The textarea's size.",
								"type":        "'small' | 'medium' | 'large'",
							},
							bson.M{
								"name":        "filled",
								"description": "Draws a filled textarea.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "label",
								"description": "The textarea's label. If you need to display HTML, use the `label` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "helpText",
								"description": "The textarea's help text. If you need to display HTML, use the `help-text` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "placeholder",
								"description": "Placeholder text to show as a hint when the input is empty.",
								"type":        "string",
							},
							bson.M{
								"name":        "rows",
								"description": "The number of rows to display by default.",
								"type":        "number",
							},
							bson.M{
								"name":        "resize",
								"description": "Controls how the textarea can be resized.",
								"type":        "'none' | 'vertical' | 'auto'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the textarea.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "readonly",
								"description": "Makes the textarea readonly.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "form",
								"description": "By default, form controls are associated with the nearest containing `<form>` element. This attribute allows you\nto place the form control outside of a form and associate it with the form that has this `id`. The form must be in\nthe same document or shadow root for this to work.",
								"type":        "string",
							},
							bson.M{
								"name":        "required",
								"description": "Makes the textarea a required field.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "minlength",
								"description": "The minimum length of input that will be considered valid.",
								"type":        "number",
							},
							bson.M{
								"name":        "maxlength",
								"description": "The maximum length of input that will be considered valid.",
								"type":        "number",
							},
							bson.M{
								"name":        "autocapitalize",
								"description": "Controls whether and how text input is automatically capitalized as it is entered by the user.",
								"type":        "'off' | 'none' | 'on' | 'sentences' | 'words' | 'characters'",
							},
							bson.M{
								"name":        "autocorrect",
								"description": "Indicates whether the browser's autocorrect feature is on or off.",
								"type":        "string",
							},
							bson.M{
								"name":        "autocomplete",
								"description": "Specifies what permission the browser has to provide assistance in filling out form field values. Refer to\nbson.A{this page on MDN}(https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) for available values.",
								"type":        "string",
							},
							bson.M{
								"name":        "autofocus",
								"description": "Indicates that the input should receive focus on page load.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "enterkeyhint",
								"description": "Used to customize the label or icon of the Enter key on virtual keyboards.",
								"type":        "'enter' | 'done' | 'go' | 'next' | 'previous' | 'search' | 'send'",
							},
							bson.M{
								"name":        "spellcheck",
								"description": "Enables spell checking on the textarea.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "inputmode",
								"description": "Tells the browser what type of data will be entered by the user, allowing it to display the appropriate virtual\nkeyboard on supportive devices.",
								"type":        "'none' | 'text' | 'decimal' | 'numeric' | 'tel' | 'search' | 'email' | 'url'",
							},
							bson.M{
								"name":        "defaultValue",
								"description": "The default value of the form control. Primarily used for resetting the form control.",
								"type":        "string",
							},
							bson.M{
								"name":        "validity",
								"description": "Gets the validity state object",
							},
							bson.M{
								"name":        "validationMessage",
								"description": "Gets the validation message",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-blur",
								"description": "Emitted when the control loses focus.",
							},
							bson.M{
								"name":        "sl-change",
								"description": "Emitted when an alteration to the control's value is committed by the user.",
							},
							bson.M{
								"name":        "sl-focus",
								"description": "Emitted when the control gains focus.",
							},
							bson.M{
								"name":        "sl-input",
								"description": "Emitted when the control receives input.",
							},
							bson.M{
								"name":        "sl-invalid",
								"description": "Emitted when the form control has been checked for validity and its constraints aren't satisfied.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-tree",
					"description": "Trees allow you to display a hierarchical list of selectable bson.A{tree items}(/components/tree-item). Items with children can be expanded and collapsed as desired by the user.\n---\n\n\n### **Events:**\n - **sl-selection-change** - Emitted when a tree item is selected or deselected.\n\n### **Slots:**\n - _default_ - The default slot.\n- **expand-icon** - The icon to show when the tree item is expanded. Works best with `<sl-icon>`.\n- **collapse-icon** - The icon to show when the tree item is collapsed. Works best with `<sl-icon>`.\n\n### **CSS Properties:**\n - **--indent-size** - The size of the indentation for nested items. _(default: var(--sl-spacing-medium))_\n- **--indent-guide-color** - The color of the indentation line. _(default: var(--sl-color-neutral-200))_\n- **--indent-guide-offset** - The amount of vertical spacing to leave between the top and bottom of the indentation line's starting position. _(default: 0)_\n- **--indent-guide-style** - The style of the indentation line, e.g. solid, dotted, dashed. _(default: solid)_\n- **--indent-guide-width** - The width of the indentation line. _(default: 0)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "selection",
							"description": "The selection behavior of the tree. Single selection allows only one node to be selected at a time. Multiple\ndisplays checkboxes and allows more than one node to be selected. Leaf allows only leaf nodes to be selected.",
							"value": bson.M{
								"type":    "'single' | 'multiple' | 'leaf'",
								"default": "'single'",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The default slot.",
						},
						bson.M{
							"name":        "expand-icon",
							"description": "The icon to show when the tree item is expanded. Works best with `<sl-icon>`.",
						},
						bson.M{
							"name":        "collapse-icon",
							"description": "The icon to show when the tree item is collapsed. Works best with `<sl-icon>`.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-selection-change",
							"type":        "bson.M{ selection: SlTreeItembson.A{} }",
							"description": "Emitted when a tree item is selected or deselected.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "expandedIconSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "collapsedIconSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name":        "selection",
								"description": "The selection behavior of the tree. Single selection allows only one node to be selected at a time. Multiple\ndisplays checkboxes and allows more than one node to be selected. Leaf allows only leaf nodes to be selected.",
								"type":        "'single' | 'multiple' | 'leaf'",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-selection-change",
								"type":        "bson.M{ selection: SlTreeItembson.A{} }",
								"description": "Emitted when a tree item is selected or deselected.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-tooltip",
					"description": "Tooltips display additional information based on a specific action.\n---\n\n\n### **Events:**\n - **sl-show** - Emitted when the tooltip begins to show.\n- **sl-after-show** - Emitted after the tooltip has shown and all animations are complete.\n- **sl-hide** - Emitted when the tooltip begins to hide.\n- **sl-after-hide** - Emitted after the tooltip has hidden and all animations are complete.\n\n### **Methods:**\n - **show()** - Shows the tooltip.\n- **hide()** - Hides the tooltip\n\n### **Slots:**\n - _default_ - The tooltip's target element. Avoid slotting in more than one element, as subsequent ones will be ignored.\n- **content** - The content to render in the tooltip. Alternatively, you can use the `content` attribute.\n\n### **CSS Properties:**\n - **--max-width** - The maximum width of the tooltip before its content will wrap. _(default: undefined)_\n- **--hide-delay** - The amount of time to wait before hiding the tooltip when hovering. _(default: undefined)_\n- **--show-delay** - The amount of time to wait before showing the tooltip when hovering. _(default: undefined)_\n\n### **CSS Parts:**\n - **base** - The component's base wrapper, an `<sl-popup>` element.\n- **base__popup** - The popup's exported `popup` part. Use this to target the tooltip's popup container.\n- **base__arrow** - The popup's exported `arrow` part. Use this to target the tooltip's arrow.\n- **body** - The tooltip's body where its content is rendered.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "content",
							"description": "The tooltip's content. If you need to display HTML, use the `content` slot instead.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "placement",
							"description": "The preferred placement of the tooltip. Note that the actual placement may vary as needed to keep the tooltip\ninside of the viewport.",
							"value": bson.M{
								"type":    "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
								"default": "'top'",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the tooltip so it won't show when triggered.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "distance",
							"description": "The distance in pixels from which to offset the tooltip away from its target.",
							"value": bson.M{
								"type":    "number",
								"default": "8",
							},
						},
						bson.M{
							"name":        "open",
							"description": "Indicates whether or not the tooltip is open. You can use this in lieu of the show/hide methods.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "skidding",
							"description": "The distance in pixels from which to offset the tooltip along its target.",
							"value": bson.M{
								"type":    "number",
								"default": "0",
							},
						},
						bson.M{
							"name":        "trigger",
							"description": "Controls how the tooltip is activated. Possible options include `click`, `hover`, `focus`, and `manual`. Multiple\noptions can be passed by separating them with a space. When manual is used, the tooltip must be activated\nprogrammatically.",
							"value": bson.M{
								"type":    "string",
								"default": "'hover focus'",
							},
						},
						bson.M{
							"name":        "hoist",
							"description": "Enable this option to prevent the tooltip from being clipped when the component is placed inside a container with\n`overflow: auto|hidden|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all,\nscenarios.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The tooltip's target element. Avoid slotting in more than one element, as subsequent ones will be ignored.",
						},
						bson.M{
							"name":        "content",
							"description": "The content to render in the tooltip. Alternatively, you can use the `content` attribute.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-show",
							"description": "Emitted when the tooltip begins to show.",
						},
						bson.M{
							"name":        "sl-after-show",
							"description": "Emitted after the tooltip has shown and all animations are complete.",
						},
						bson.M{
							"name":        "sl-hide",
							"description": "Emitted when the tooltip begins to hide.",
						},
						bson.M{
							"name":        "sl-after-hide",
							"description": "Emitted after the tooltip has hidden and all animations are complete.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "body",
								"type": "HTMLElement",
							},
							bson.M{
								"name": "popup",
								"type": "SlPopup",
							},
							bson.M{
								"name":        "content",
								"description": "The tooltip's content. If you need to display HTML, use the `content` slot instead.",
								"type":        "string",
							},
							bson.M{
								"name":        "placement",
								"description": "The preferred placement of the tooltip. Note that the actual placement may vary as needed to keep the tooltip\ninside of the viewport.",
								"type":        "| 'top'\n    | 'top-start'\n    | 'top-end'\n    | 'right'\n    | 'right-start'\n    | 'right-end'\n    | 'bottom'\n    | 'bottom-start'\n    | 'bottom-end'\n    | 'left'\n    | 'left-start'\n    | 'left-end'",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the tooltip so it won't show when triggered.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "distance",
								"description": "The distance in pixels from which to offset the tooltip away from its target.",
								"type":        "number",
							},
							bson.M{
								"name":        "open",
								"description": "Indicates whether or not the tooltip is open. You can use this in lieu of the show/hide methods.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "skidding",
								"description": "The distance in pixels from which to offset the tooltip along its target.",
								"type":        "number",
							},
							bson.M{
								"name":        "trigger",
								"description": "Controls how the tooltip is activated. Possible options include `click`, `hover`, `focus`, and `manual`. Multiple\noptions can be passed by separating them with a space. When manual is used, the tooltip must be activated\nprogrammatically.",
								"type":        "string",
							},
							bson.M{
								"name":        "hoist",
								"description": "Enable this option to prevent the tooltip from being clipped when the component is placed inside a container with\n`overflow: auto|hidden|scroll`. Hoisting uses a fixed positioning strategy that works in many, but not all,\nscenarios.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-show",
								"description": "Emitted when the tooltip begins to show.",
							},
							bson.M{
								"name":        "sl-after-show",
								"description": "Emitted after the tooltip has shown and all animations are complete.",
							},
							bson.M{
								"name":        "sl-hide",
								"description": "Emitted when the tooltip begins to hide.",
							},
							bson.M{
								"name":        "sl-after-hide",
								"description": "Emitted after the tooltip has hidden and all animations are complete.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-tree-item",
					"description": "A tree item serves as a hierarchical node that lives inside a bson.A{tree}(/components/tree).\n---\n\n\n### **Events:**\n - **sl-expand** - Emitted when the tree item expands.\n- **sl-after-expand** - Emitted after the tree item expands and all animations are complete.\n- **sl-collapse** - Emitted when the tree item collapses.\n- **sl-after-collapse** - Emitted after the tree item collapses and all animations are complete.\n- **sl-lazy-change** - Emitted when the tree item's lazy state changes.\n- **sl-lazy-load** - Emitted when a lazy item is selected. Use this event to asynchronously load data and append items to the tree before expanding. After appending new items, remove the `lazy` attribute to remove the loading state and update the tree.\n\n### **Methods:**\n - **getChildrenItems(bson.M{ includeDisabled = true }: _bson.M{ includeDisabled?: boolean }_): _SlTreeItembson.A{}_** - Gets all the nested tree items in this node.\n\n### **Slots:**\n - _default_ - The default slot.\n- **expand-icon** - The icon to show when the tree item is expanded.\n- **collapse-icon** - The icon to show when the tree item is collapsed.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **item** - The tree item's container. This element wraps everything except slotted tree item children.\n- **item--disabled** - Applied when the tree item is disabled.\n- **item--expanded** - Applied when the tree item is expanded.\n- **item--indeterminate** - Applied when the selection is indeterminate.\n- **item--selected** - Applied when the tree item is selected.\n- **indentation** - The tree item's indentation container.\n- **expand-button** - The container that wraps the tree item's expand button and spinner.\n- **spinner** - The spinner that shows when a lazy tree item is in the loading state.\n- **spinner__base** - The spinner's base part.\n- **label** - The tree item's label.\n- **children** - The container that wraps the tree item's nested children.\n- **checkbox** - The checkbox that shows when using multiselect.\n- **checkbox__base** - The checkbox's exported `base` part.\n- **checkbox__control** - The checkbox's exported `control` part.\n- **checkbox__control--checked** - The checkbox's exported `control--checked` part.\n- **checkbox__control--indeterminate** - The checkbox's exported `control--indeterminate` part.\n- **checkbox__checked-icon** - The checkbox's exported `checked-icon` part.\n- **checkbox__indeterminate-icon** - The checkbox's exported `indeterminate-icon` part.\n- **checkbox__label** - The checkbox's exported `label` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "expanded",
							"description": "Expands the tree item.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "selected",
							"description": "Draws the tree item in a selected state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the tree item.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "lazy",
							"description": "Enables lazy loading behavior.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The default slot.",
						},
						bson.M{
							"name":        "expand-icon",
							"description": "The icon to show when the tree item is expanded.",
						},
						bson.M{
							"name":        "collapse-icon",
							"description": "The icon to show when the tree item is collapsed.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-expand",
							"description": "Emitted when the tree item expands.",
						},
						bson.M{
							"name":        "sl-after-expand",
							"description": "Emitted after the tree item expands and all animations are complete.",
						},
						bson.M{
							"name":        "sl-collapse",
							"description": "Emitted when the tree item collapses.",
						},
						bson.M{
							"name":        "sl-after-collapse",
							"description": "Emitted after the tree item collapses and all animations are complete.",
						},
						bson.M{
							"name":        "sl-lazy-change",
							"description": "Emitted when the tree item's lazy state changes.",
						},
						bson.M{
							"name":        "sl-lazy-load",
							"description": "Emitted when a lazy item is selected. Use this event to asynchronously load data and append items to the tree before expanding. After appending new items, remove the `lazy` attribute to remove the loading state and update the tree.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "indeterminate",
								"type": "boolean",
							},
							bson.M{
								"name": "isLeaf",
								"type": "boolean",
							},
							bson.M{
								"name": "loading",
								"type": "boolean",
							},
							bson.M{
								"name": "selectable",
								"type": "boolean",
							},
							bson.M{
								"name":        "expanded",
								"description": "Expands the tree item.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "selected",
								"description": "Draws the tree item in a selected state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the tree item.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "lazy",
								"description": "Enables lazy loading behavior.",
								"type":        "boolean",
							},
							bson.M{
								"name": "defaultSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "childrenSlot",
								"type": "HTMLSlotElement",
							},
							bson.M{
								"name": "itemElement",
								"type": "HTMLDivElement",
							},
							bson.M{
								"name": "childrenContainer",
								"type": "HTMLDivElement",
							},
							bson.M{
								"name": "expandButtonSlot",
								"type": "HTMLSlotElement",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-expand",
								"description": "Emitted when the tree item expands.",
							},
							bson.M{
								"name":        "sl-after-expand",
								"description": "Emitted after the tree item expands and all animations are complete.",
							},
							bson.M{
								"name":        "sl-collapse",
								"description": "Emitted when the tree item collapses.",
							},
							bson.M{
								"name":        "sl-after-collapse",
								"description": "Emitted after the tree item collapses and all animations are complete.",
							},
							bson.M{
								"name":        "sl-lazy-change",
								"description": "Emitted when the tree item's lazy state changes.",
							},
							bson.M{
								"name":        "sl-lazy-load",
								"description": "Emitted when a lazy item is selected. Use this event to asynchronously load data and append items to the tree before expanding. After appending new items, remove the `lazy` attribute to remove the loading state and update the tree.",
							},
						},
					},
				},
				bson.M{
					"name":        "sl-visually-hidden",
					"description": "The visually hidden utility makes content accessible to assistive devices without displaying it on the screen.\n---\n\n\n### **Slots:**\n - _default_ - The content to be visually hidden.",
					"doc-url":     "",
					"attributes":  bson.A{},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The content to be visually hidden.",
						},
					},
					"events": bson.A{},
					"js": bson.M{
						"properties": bson.A{},
						"events":     bson.A{},
					},
				},
				bson.M{
					"name":        "sl-tab",
					"description": "Tabs are used inside bson.A{tab groups}(/components/tab-group) to represent and activate bson.A{tab panels}(/components/tab-panel).\n---\n\n\n### **Events:**\n - **sl-close** - Emitted when the tab is closable and the close button is activated.\n\n### **Slots:**\n - _default_ - The tab's label.\n\n### **CSS Parts:**\n - **base** - The component's base wrapper.\n- **close-button** - The close button, an `<sl-icon-button>`.\n- **close-button__base** - The close button's exported `base` part.",
					"doc-url":     "",
					"attributes": bson.A{
						bson.M{
							"name":        "panel",
							"description": "The name of the tab panel this tab is associated with. The panel must be located in the same tab group.",
							"value": bson.M{
								"type":    "string",
								"default": "''",
							},
						},
						bson.M{
							"name":        "active",
							"description": "Draws the tab in an active state.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "closable",
							"description": "Makes the tab closable and shows a close button.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
						bson.M{
							"name":        "disabled",
							"description": "Disables the tab and prevents selection.",
							"value": bson.M{
								"type":    "boolean",
								"default": "false",
							},
						},
					},
					"slots": bson.A{
						bson.M{
							"name":        "",
							"description": "The tab's label.",
						},
					},
					"events": bson.A{
						bson.M{
							"name":        "sl-close",
							"description": "Emitted when the tab is closable and the close button is activated.",
						},
					},
					"js": bson.M{
						"properties": bson.A{
							bson.M{
								"name": "tab",
								"type": "HTMLElement",
							},
							bson.M{
								"name":        "panel",
								"description": "The name of the tab panel this tab is associated with. The panel must be located in the same tab group.",
								"type":        "string",
							},
							bson.M{
								"name":        "active",
								"description": "Draws the tab in an active state.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "closable",
								"description": "Makes the tab closable and shows a close button.",
								"type":        "boolean",
							},
							bson.M{
								"name":        "disabled",
								"description": "Disables the tab and prevents selection.",
								"type":        "boolean",
							},
						},
						"events": bson.A{
							bson.M{
								"name":        "sl-close",
								"description": "Emitted when the tab is closable and the close button is activated.",
							},
						},
					},
				},
			},
		},
	},
}
