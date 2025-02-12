package llm

var editRules = `Here are the rules that describe how to approach the editing process:
- Decide what you’re actually saying.
- Who are you writing for?
- What is your main point? 

Repeat yourself (within reason)
- Look for ways that you can restate your point, clarify, or provide closure for the reader.

Simplify
- Strive to get to the point as quickly as possible.

Eliminate passive voice
- Rewrite a passive construction to active to make what you’re saying clearer and make the sentence easier to read.

Don’t use adverbs
- Replace an adverb with a better, more specific verb, or describe what you mean instead.

Don’t assume knowledge
- Spell out acronyms on first use.

Be aware of your tone
- Know what kind of tone you’re going for, and be consistent.
- Don't switch between formal / non-formal tones.`

var contentRules = `Here are the rules that describe how to handle the content:
- NEVER change the contents of a code block.
- NEVER change links.
- Avoid introducing any errors to the content structure (e.g. Yaml header).`

var outputRules = `Here are the rules that describe how to output the changes:
- For each optimization you find, output a separate change that can be replaced directly in the text.
- Each search term MUST be unique and match exactly once.
- Always consider the context of the full sentence before making a change and do not introduce grammatical errors.`
