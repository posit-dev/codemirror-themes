function materialDarkHighlightStyle() {
  const config = {
    name: 'materialDark',
    dark: true,
    background: '#263238',
    foreground: '#EEFFFF',
    selection: '#80CBC420',
    cursor: '#FFCC00',
    dropdownBackground: '#263238',
    dropdownBorder: '#FFFFFF10',
    activeLine: '#00000050',
    matchingBracket: '#263238',
    keyword: '#89DDFF',
    storage: '#89DDFF',
    variable: '#EEFFFF',
    parameter: '#EEFFFF',
    function: '#82AAFF',
    string: '#C3E88D',
    constant: '#89DDFF',
    type: '#FFCB6B',
    class: '#FFCB6B',
    number: '#F78C6C',
    comment: '#546E7A',
    heading: '#89DDFF',
    invalid: '#f0717870',
    regexp: '#C3E88D',
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
