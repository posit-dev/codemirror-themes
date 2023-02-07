function githubDarkHighlightStyle() {
  const config = {
    name: 'githubDark',
    dark: true,
    background: '#24292e',
    foreground: '#d1d5da',
    selection: '#3392FF44',
    cursor: '#c8e1ff',
    dropdownBackground: '#24292e',
    dropdownBorder: '#1b1f23',
    activeLine: '#2b3036',
    matchingBracket: '#17E5E650',
    keyword: '#f97583',
    storage: '#f97583',
    variable: '#ffab70',
    parameter: '#e1e4e8',
    function: '#79b8ff',
    string: '#9ecbff',
    constant: '#79b8ff',
    type: '#79b8ff',
    class: '#b392f0',
    number: '#79b8ff',
    comment: '#6a737d',
    heading: '#79b8ff',
    invalid: '#f97583',
    regexp: '#9ecbff',
  };
  return HighlightStyle.define([ 
    {tag: t.keyword, color: config.keyword},
    {tag: [t.name, t.deleted, t.character, t.macroName], color: config.variable},
    {tag: [t.propertyName], color: config.function},
    {tag: [t.processingInstruction, t.string, t.inserted, t.special(t.string)], color: config.string},
    {tag: [t.function(t.variableName), t.labelName], color: config.function},
    {tag: [t.color, t.constant(t.name), t.standard(t.name)], color: config.constant},
    {tag: [t.definition(t.name), t.separator], color: config.variable},
    {tag: [t.className], color: config.class},
    {tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace], color: config.number},
    {tag: [t.typeName], color: config.type, fontStyle: config.type},
    {tag: [t.operator, t.operatorKeyword], color: config.keyword},
    {tag: [t.url, t.escape, t.regexp, t.link], color: config.regexp},
    {tag: [t.meta, t.comment], color: config.comment},
    {tag: t.strong, fontWeight: 'bold'},
    {tag: t.emphasis, fontStyle: 'italic'},
    {tag: t.link, textDecoration: 'underline'},
    {tag: t.heading, fontWeight: 'bold', color: config.heading},
    {tag: [t.atom, t.bool, t.special(t.variableName)], color: config.variable},
    {tag: t.invalid, color: config.invalid},
    {tag: t.strikethrough, textDecoration: 'line-through'},
  ]);
}
