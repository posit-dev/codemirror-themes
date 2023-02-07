function monokaiHighlightStyle() {
  const config = {
    name: 'monokai',
    dark: true,
    background: '#272822',
    foreground: '#f8f8f2',
    selection: '#878b9180',
    cursor: '#f8f8f0',
    dropdownBackground: '#272822',
    dropdownBorder: '#75715E',
    activeLine: '#3e3d32',
    matchingBracket: '#3e3d32',
    keyword: '#F92672',
    storage: '#F92672',
    variable: '#F8F8F2',
    parameter: '#FD971F',
    function: '#A6E22E',
    string: '#E6DB74',
    constant: '#AE81FF',
    type: '#A6E22E',
    class: '#A6E22E',
    number: '#AE81FF',
    comment: '#88846f',
    heading: '#A6E22E',
    invalid: '#F44747',
    regexp: '#E6DB74',
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
